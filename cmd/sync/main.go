// Command sync перестраивает NEETCODE_PLAN_*.md из двух источников:
//   - catalog.yaml (статичные метаданные роадмапа: id, название, сложность, паттерн, блок, группа)
//   - Obsidian vault (прогресс: solved_dates, needRepeat, retention, algorithms, created)
//
// Связывание по номеру LeetCode в префиксе имени заметки ("345. ...md" → 345).
//
// Примеры:
//
//	rtk go run ./cmd/sync --init          # первичный импорт планов в catalog.yaml
//	rtk go run ./cmd/sync                 # перегенерировать планы из каталога + vault
//	rtk go run ./cmd/sync --check         # только диагностика, без записи
//	rtk go run ./cmd/sync --dry-run       # показать сводный diff, не писать
package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// -------------------- Типы каталога --------------------

// Catalog — корневая структура catalog.yaml.
type Catalog struct {
	Version   int     `yaml:"version"`
	VaultRoot string  `yaml:"vault_root"`
	Blocks    []Block `yaml:"blocks"`
}

// Block — блок плана (один NEETCODE_PLAN_*.md).
type Block struct {
	ID       string    `yaml:"id"`
	Title    string    `yaml:"title"`
	Folder   string    `yaml:"folder"`    // имя папки внутри vault_root
	PlanFile string    `yaml:"plan_file"` // имя выходного файла
	Sections []Section `yaml:"sections"`  // опционально: порядок и заголовки групп
	Problems []Problem `yaml:"problems"`
}

// Section — опциональное описание секции (группа + заголовок).
type Section struct {
	Group string `yaml:"group"`
	Title string `yaml:"title"`
}

// Problem — статичные метаданные задачи.
type Problem struct {
	LeetcodeID int    `yaml:"leetcode_id"`
	Title      string `yaml:"title"`
	Difficulty string `yaml:"difficulty"` // Easy | Medium | Hard
	Pattern    string `yaml:"pattern"`
	Group      string `yaml:"group"` // main | bonus | yandex | extra
}

// -------------------- Типы прогресса из Obsidian --------------------

// NoteProgress — распарсенный frontmatter заметки.
type NoteProgress struct {
	LeetcodeID  int      // извлекается из имени файла
	Title       string   // имя файла без расширения и префикса-номера
	Folder      string   // подпапка внутри vault (имя блока по соглашению)
	Path        string   // полный путь к заметке
	Created     string   `yaml:"created"`
	SolvedDates []string `yaml:"solved_dates"`
	NeedRepeat  bool     `yaml:"needRepeat"`
	Retention   flexInt  `yaml:"retention"` // 0 = не проставлено
	Difficulty  string   `yaml:"difficulty"`
	Algorithms  []string `yaml:"algorithms"`
}

// flexInt принимает как int (3), так и строку ("3") или ("L3").
// Защищает от падения парсера, если Obsidian-поле записано в кавычках.
type flexInt int

// UnmarshalYAML реализует кастомный парсинг flexInt.
func (f *flexInt) UnmarshalYAML(value *yaml.Node) error {
	var n int
	if err := value.Decode(&n); err == nil {
		*f = flexInt(n)
		return nil
	}
	// Падает как int — пробуем как строку и выцепляем число.
	var s string
	if err := value.Decode(&s); err != nil {
		return err
	}
	// Извлекаем первую последовательность цифр (для "3", "L3", "level 3").
	m := flexIntRe.FindString(s)
	if m == "" {
		*f = 0
		return nil
	}
	v, err := strconv.Atoi(m)
	if err != nil {
		return err
	}
	*f = flexInt(v)
	return nil
}

// flexIntRe — последовательность цифр в строке.
var flexIntRe = regexp.MustCompile(`\d+`)

// -------------------- Загрузка каталога --------------------

// LoadCatalog читает и валидирует catalog.yaml.
func LoadCatalog(path string) (*Catalog, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read catalog: %w", err)
	}
	var c Catalog
	if err := yaml.Unmarshal(data, &c); err != nil {
		return nil, fmt.Errorf("parse catalog yaml: %w", err)
	}
	if err := c.Validate(); err != nil {
		return nil, err
	}
	return &c, nil
}

// Validate проверяет целостность каталога.
func (c *Catalog) Validate() error {
	if c.VaultRoot == "" {
		return errors.New("catalog: vault_root is empty")
	}
	if len(c.Blocks) == 0 {
		return errors.New("catalog: no blocks defined")
	}
	seenFolder := make(map[string]string)
	for _, b := range c.Blocks {
		if b.ID == "" {
			return errors.New("catalog: block without id")
		}
		if b.PlanFile == "" {
			return fmt.Errorf("catalog: block %q has no plan_file", b.ID)
		}
		if b.Folder == "" {
			return fmt.Errorf("catalog: block %q has no folder", b.ID)
		}
		if _, dup := seenFolder[b.Folder]; dup {
			return fmt.Errorf("catalog: folder %q used by blocks %q and %q", b.Folder, seenFolder[b.Folder], b.ID)
		}
		seenFolder[b.Folder] = b.ID
		// Дубль leetcode_id внутри одного блока — ошибка (одна задача дважды в блоке).
		// Между блоками — разрешён (задача может изучаться разными паттернами).
		seenInBlock := make(map[int]bool)
		for _, p := range b.Problems {
			if p.LeetcodeID == 0 {
				return fmt.Errorf("catalog: block %q has problem without leetcode_id (%q)", b.ID, p.Title)
			}
			if seenInBlock[p.LeetcodeID] {
				return fmt.Errorf("catalog: leetcode_id %d duplicated within block %q", p.LeetcodeID, b.ID)
			}
			seenInBlock[p.LeetcodeID] = true
			switch p.Difficulty {
			case "", "Easy", "Medium", "Hard":
			default:
				return fmt.Errorf("catalog: problem %d has invalid difficulty %q", p.LeetcodeID, p.Difficulty)
			}
		}
	}
	return nil
}

// FindBlockByFolder возвращает блок по имени папки vault.
func (c *Catalog) FindBlockByFolder(folder string) *Block {
	for i := range c.Blocks {
		if c.Blocks[i].Folder == folder {
			return &c.Blocks[i]
		}
	}
	return nil
}

// -------------------- Сканирование vault --------------------

var (
	// leetcodeIDRe — номер в начале имени файла: "345. Reverse Vowels..." → 345.
	leetcodeIDRe = regexp.MustCompile(`^(\d+)\.\s+(.+)$`)
	// frontmatterDelim — строка из трёх дефисов.
	frontmatterDelim = "---"
)

// ScanVault обходит vault_root и собирает прогресс из всех заметок.
// Возвращает карту leetcode_id → прогресс и список предупреждений.
func ScanVault(root string) (map[int]*NoteProgress, []string, error) {
	progress := make(map[int]*NoteProgress)
	var warnings []string
	var dupIDs []int

	walkErr := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".md") {
			return nil
		}
		np, w, err := parseNote(root, path)
		if err != nil {
			warnings = append(warnings, err.Error())
			return nil
		}
		warnings = append(warnings, w...)
		if np == nil {
			return nil
		}
		if _, exists := progress[np.LeetcodeID]; exists {
			dupIDs = append(dupIDs, np.LeetcodeID)
		}
		progress[np.LeetcodeID] = np
		return nil
	})
	if walkErr != nil {
		return nil, warnings, fmt.Errorf("walk vault: %w", walkErr)
	}
	sort.Ints(dupIDs)
	for _, id := range dupIDs {
		warnings = append(warnings, fmt.Sprintf("⚠️  несколько заметок с leetcode_id %d (взята последняя)", id))
	}
	return progress, warnings, nil
}

// parseNote читает одну заметку, извлекает frontmatter и leetcode_id.
func parseNote(vaultRoot, path string) (*NoteProgress, []string, error) {
	var warnings []string
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, fmt.Errorf("open note %s: %w", path, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("scan note %s: %w", path, err)
	}

	// Извлечение frontmatter: блок между первой и второй строкой "---".
	fm, ok := extractFrontmatter(lines)
	if !ok {
		// Нет frontmatter — не ошибка для Obsidian, но пропускаем как не-задачу.
		return nil, nil, nil
	}

	// Фильтр: только Type: exercise и tags содержит algorithms.
	if !isExerciseNote(fm) {
		return nil, nil, nil
	}

	np := &NoteProgress{Path: path}
	if err := yaml.Unmarshal([]byte(fm), np); err != nil {
		return nil, nil, fmt.Errorf("parse frontmatter %s: %w", path, err)
	}

	base := filepath.Base(path)
	name := strings.TrimSuffix(base, ".md")
	if m := leetcodeIDRe.FindStringSubmatch(name); m != nil {
		id, err := strconv.Atoi(m[1])
		if err != nil {
			return nil, nil, fmt.Errorf("parse leetcode_id from %q: %w", name, err)
		}
		np.LeetcodeID = id
		np.Title = m[2]
	} else {
		// Нет номера в имени — не алгоритм-задача или другая схема.
		warnings = append(warnings, fmt.Sprintf("⚠️  заметка без leetcode_id в имени: %s", base))
		return nil, warnings, nil
	}

	rel, _ := filepath.Rel(vaultRoot, path)
	np.Folder = filepath.Dir(rel)
	if np.Folder == "." {
		np.Folder = ""
	}
	return np, warnings, nil
}

// extractFrontmatter возвращает текст YAML между маркерами "---".
func extractFrontmatter(lines []string) (string, bool) {
	if len(lines) == 0 || strings.TrimSpace(lines[0]) != frontmatterDelim {
		return "", false
	}
	var b strings.Builder
	for i := 1; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == frontmatterDelim {
			return b.String(), true
		}
		b.WriteString(lines[i])
		b.WriteByte('\n')
	}
	return "", false // нет закрывающего "---"
}

// isExerciseNote проверяет по сыром YAML, что Type содержит exercise и есть tag algorithms/leetcode.
func isExerciseNote(fm string) bool {
	// Грубая, но дешёвая проверка по строкам — избегаем полного unmarshal дважды.
	hasExercise := false
	hasAlgoTag := false
	for _, line := range strings.Split(fm, "\n") {
		trim := strings.TrimSpace(line)
		if strings.Contains(trim, "exercise") {
			hasExercise = true
		}
		if strings.Contains(trim, "algorithms") {
			hasAlgoTag = true
		}
	}
	return hasExercise && hasAlgoTag
}

// -------------------- Слияние и рендер --------------------

// retentionEmoji переводит числовую оценку в эмодзи-метку.
func retentionEmoji(r flexInt) string {
	switch r {
	case 5:
		return "⚪ L5"
	case 4:
		return "🟢 L4"
	case 3:
		return "🟡 L3"
	case 2:
		return "🟠 L2"
	case 1:
		return "🔴 L1"
	default:
		return ""
	}
}

// needRepeatCell возвращает "true" (без кавычек), если задача требует повторения.
func needRepeatCell(np *NoteProgress) string {
	if np != nil && np.NeedRepeat {
		return "true"
	}
	return ""
}

// statusMark возвращает [x] для решённой, [ ] для нерешённой.
func statusMark(np *NoteProgress) string {
	if np != nil {
		return "[x]"
	}
	return "[ ]"
}

// solvedCount возвращает количество решений (len(solved_dates)) как строку.
func solvedCount(np *NoteProgress) string {
	if np == nil || len(np.SolvedDates) == 0 {
		return ""
	}
	return strconv.Itoa(len(np.SolvedDates))
}

// firstSolvedDate возвращает первую (раннюю) дату из solved_dates.
func firstSolvedDate(np *NoteProgress) string {
	if np == nil || len(np.SolvedDates) == 0 {
		return ""
	}
	dates := append([]string(nil), np.SolvedDates...)
	sort.Strings(dates)
	return dates[0]
}

// defaultGroupTitles — заголовки секций по умолчанию, если блок не задал Sections.
var defaultGroupTitles = map[string]string{
	"main":   "Основной блок",
	"bonus":  "Bonus",
	"yandex": "Дополнительная практика: Задачи с собеседований Яндекс",
	"extra":  "Вне каталога (Extra)",
}

// sectionTitle возвращает заголовок секции для группы блока.
func (b *Block) sectionTitle(group string) string {
	for _, s := range b.Sections {
		if s.Group == group {
			return s.Title
		}
	}
	if t, ok := defaultGroupTitles[group]; ok {
		return t
	}
	return group
}

// orderedGroups возвращает группы в порядке: Sections, затем остальные по discovery.
func (b *Block) orderedGroups(present map[string]bool) []string {
	var out []string
	seen := make(map[string]bool)
	for _, s := range b.Sections {
		if present[s.Group] && !seen[s.Group] {
			out = append(out, s.Group)
			seen[s.Group] = true
		}
	}
	// Гарантированный порядок для стандартных групп.
	for _, g := range []string{"main", "bonus", "yandex", "extra"} {
		if present[g] && !seen[g] {
			out = append(out, g)
			seen[g] = true
		}
	}
	// Любые нестандартные группы — по алфавиту.
	var rest []string
	for g := range present {
		if !seen[g] {
			rest = append(rest, g)
		}
	}
	sort.Strings(rest)
	out = append(out, rest...)
	return out
}

// rowTaskCell — колонка «Задача»: bold для существующих.
func rowTaskCell(p Problem) string {
	return fmt.Sprintf("**%d. %s**", p.LeetcodeID, p.Title)
}

// rowTaskCellOrphan — для orphan: только заголовок из заметки.
func rowTaskCellOrphan(np *NoteProgress) string {
	return fmt.Sprintf("**%d. %s**", np.LeetcodeID, np.Title)
}

// renderRow — одна строка таблицы.
func renderRow(p Problem, np *NoteProgress) string {
	status := statusMark(np)
	task := rowTaskCell(p)
	diff := p.Difficulty
	pattern := p.Pattern
	date := firstSolvedDate(np)
	ret := ""
	if np != nil {
		ret = retentionEmoji(np.Retention)
	}
	nr := needRepeatCell(np)
	count := solvedCount(np)
	return fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s | %s |",
		status, task, diff, pattern, date, ret, nr, count)
}

// renderOrphanRow — строка для orphan-заметки (без записи в каталоге).
func renderOrphanRow(np *NoteProgress) string {
	status := statusMark(np)
	task := rowTaskCellOrphan(np)
	pattern := strings.Join(np.Algorithms, ", ")
	date := firstSolvedDate(np)
	ret := retentionEmoji(np.Retention)
	nr := needRepeatCell(np)
	count := solvedCount(np)
	return fmt.Sprintf("| %s | %s |  | %s | %s | %s | %s | %s |",
		status, task, pattern, date, ret, nr, count)
}

// planHeader — шапка плана (легенда + правило + разминка + как работаем).
func planHeader(blockTitle string) string {
	return fmt.Sprintf(`# NeetCode Roadmap — %s

> **Сгенерировано командой `+"`rtk go run ./cmd/sync`"+`.**
> Не редактируй вручную — правь catalog.yaml (статика) и заметки Obsidian (прогресс).

---

## Легенда оценки (Retention)

| Метка | Значение | Когда ставить |
|-------|----------|---------------|
| 🟢 **L4** | Решила сама, быстро, чистый код | Без подсказок, за разумное время, сразу верный Big O |
| 🟡 **L3** | Решила сама, но медленно или с багами | Сама, но долго думала или баги в edge cases |
| 🟠 **L2** | Нужны были подсказки по алгоритму | Я направил к паттерну, но код написала сама |
| 🔴 **L1** | Существенная помощь, нужно повторить | Помогал с алгоритмом и/или кодом |
| ⚪ **L5** | Эксперт — могу объяснить другому | Решила, оптимизировала, знаю все edge cases |

- **Статус**: `+"`[x]`"+` — решена (есть заметка в Obsidian), `+"`[ ]`"+` — ещё не решена.
- **needRepeat**: `+"`true`"+` — в Obsidian стоит флаг «нужно повторить».
- **Решений**: количество дат в `+"`solved_dates`"+` (сколько раз решалась).

---

## Правило перехода

> **Не начинать новую задачу, пока текущая не доведена до L4.**
>
> Исключение: разминка (warm-up) перед новой задачей — повторение уже решённой задачи.

---

## Ежедневная разминка (Repetition)

0. **Warm-up**: Перед новой задачей решаем одну из уже выполненных самостоятельно.
   - **Вербализация**: Сначала описываешь алгоритм и код словами (почему и как).
   - **Zero Hints**: Без подсказок во время решения разминки.
   - **Без тестов**: Моделируем ситуацию на собеседовании.

---

## Как работаем

1. **TDD**: Сначала тесты в `+"`day_X/main_test.go`"+`.
2. **Go Way**: Используем `+"`make`"+` с аллокацией, `+"`struct{}`"+` для сетов, `+"`range`"+` для итераций.
3. **Complexity**: Анализируем Time и Space Complexity для каждого решения.

---
`, blockTitle)
}

// tableHeader — заголовок таблицы плана.
func tableHeader() string {
	return "| Статус | Задача | Сложность | Тема / Паттерн | Дата | Retention | needRepeat | Решений |\n| :---: | :--- | :---: | :--- | :--- | :--- | :---: | :---: |"
}

// RenderPlan строит содержимое NEETCODE_PLAN_<n>.md для блока.
func RenderPlan(b *Block, progress map[int]*NoteProgress) string {
	var out bytes.Buffer
	out.WriteString(planHeader(b.Title))

	// Группируем задачи каталога по группам.
	byGroup := make(map[string][]Problem)
	for _, p := range b.Problems {
		g := p.Group
		if g == "" {
			g = "main"
		}
		byGroup[g] = append(byGroup[g], p)
	}

	// Собираем orphan-заметки, относящиеся к этому блоку по папке.
	var orphans []*NoteProgress
	for _, np := range progress {
		if _, inCatalog := findProblem(b, np.LeetcodeID); inCatalog {
			continue
		}
		if folderBlock(np.Folder) == b.Folder || np.Folder == b.Folder {
			orphans = append(orphans, np)
		}
	}
	sort.Slice(orphans, func(i, j int) bool {
		return orphans[i].LeetcodeID < orphans[j].LeetcodeID
	})
	if len(orphans) > 0 {
		byGroup["extra"] = nil // группа присутствует
	}

	// Определяем присутствующие группы.
	present := make(map[string]bool)
	for g := range byGroup {
		present[g] = true
	}
	for _, p := range b.Problems {
		g := p.Group
		if g == "" {
			g = "main"
		}
		present[g] = true
	}
	if len(orphans) > 0 {
		present["extra"] = true
	}

	// Рендерим секции в порядке.
	for _, g := range b.orderedGroups(present) {
		title := b.sectionTitle(g)
		fmt.Fprintf(&out, "## %s\n\n", title)
		out.WriteString(tableHeader())
		out.WriteByte('\n')

		switch g {
		case "extra":
			for _, np := range orphans {
				out.WriteString(renderOrphanRow(np))
				out.WriteByte('\n')
			}
		default:
			ps := byGroup[g]
			sort.Slice(ps, func(i, j int) bool {
				return ps[i].LeetcodeID < ps[j].LeetcodeID
			})
			for _, p := range ps {
				np := progress[p.LeetcodeID]
				out.WriteString(renderRow(p, np))
				out.WriteByte('\n')
			}
		}
		out.WriteByte('\n')
	}

	// Подсказка для orphan.
	if len(orphans) > 0 {
		out.WriteString("> 📌 Заметки выше решены, но отсутствуют в `catalog.yaml`. ")
		out.WriteString("Если они часть роадмапа — добавь их в каталог.\n\n")
	}

	return out.String()
}

// findProblem ищет задачу по leetcode_id в блоке.
func findProblem(b *Block, id int) (Problem, bool) {
	for _, p := range b.Problems {
		if p.LeetcodeID == id {
			return p, true
		}
	}
	return Problem{}, false
}

// folderBlock возвращает верхнеуровневую подпапку (первый сегмент).
func folderBlock(folder string) string {
	if folder == "" || folder == "." {
		return ""
	}
	if idx := strings.Index(folder, string(os.PathSeparator)); idx >= 0 {
		return folder[:idx]
	}
	return folder
}

// ImportOrphans добавляет orphan-заметки из progress в соответствующие блоки каталога.
// Возвращает количество добавленных задач.
func ImportOrphans(cat *Catalog, progress map[int]*NoteProgress) int {
	existing := make(map[int]bool)
	for _, b := range cat.Blocks {
		for _, p := range b.Problems {
			existing[p.LeetcodeID] = true
		}
	}

	var ids []int
	for id := range progress {
		ids = append(ids, id)
	}
	sort.Ints(ids)

	added := 0
	for _, id := range ids {
		if existing[id] {
			continue
		}
		np := progress[id]
		topFolder := folderBlock(np.Folder)
		b := cat.FindBlockByFolder(topFolder)
		if b == nil {
			b = cat.FindBlockByFolder(np.Folder)
		}
		if b == nil {
			continue
		}

		diff := strings.TrimSpace(np.Difficulty)
		switch strings.ToLower(diff) {
		case "easy":
			diff = "Easy"
		case "medium":
			diff = "Medium"
		case "hard":
			diff = "Hard"
		default:
			diff = "Easy"
		}

		pattern := strings.Join(np.Algorithms, ", ")

		p := Problem{
			LeetcodeID: np.LeetcodeID,
			Title:      np.Title,
			Difficulty: diff,
			Pattern:    pattern,
			Group:      "main",
		}
		b.Problems = append(b.Problems, p)
		existing[id] = true
		added++
	}

	return added
}

// -------------------- Запуск --------------------

type runOptions struct {
	catalogPath string
	vaultRoot   string
	checkOnly   bool
	dryRun      bool
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "sync: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	fs := flag.NewFlagSet("sync", flag.ContinueOnError)
	initMode := fs.Bool("init", false, "первичный импорт текущих NEETCODE_PLAN_*.md в catalog.yaml")
	check := fs.Bool("check", false, "только диагностика, без записи файлов")
	dryRun := fs.Bool("dry-run", false, "показать планируемые изменения, не записывая")
	catalogPath := fs.String("catalog", "catalog.yaml", "путь к catalog.yaml")
	vaultFlag := fs.String("vault", "", "путь к vault (по умолчанию из catalog.yaml)")
	importOrphansFlag := fs.Bool("import-orphans", false, "импортировать orphan-заметки из vault в catalog.yaml")
	plansDir := fs.String("plans-dir", ".", "директория с текущими NEETCODE_PLAN_*.md (для --init)")
	outDir := fs.String("out-dir", ".", "директория для записи NEETCODE_PLAN_*.md")
	if err := fs.Parse(args); err != nil {
		return err
	}

	if *initMode {
		return runInit(*plansDir, *catalogPath)
	}

	cat, err := LoadCatalog(*catalogPath)
	if err != nil {
		return err
	}
	vaultRoot := *vaultFlag
	if vaultRoot == "" {
		vaultRoot = cat.VaultRoot
	}

	progress, warns, err := ScanVault(vaultRoot)
	if err != nil {
		return err
	}

	if *importOrphansFlag {
		added := ImportOrphans(cat, progress)
		if added > 0 {
			if err := cat.Validate(); err != nil {
				return fmt.Errorf("validate catalog after orphan import: %w", err)
			}
			out, err := yaml.Marshal(cat)
			if err != nil {
				return fmt.Errorf("marshal catalog: %w", err)
			}
			if *dryRun {
				fmt.Printf("[dry-run] %s: добавлено %d задач\n", *catalogPath, added)
			} else {
				if err := os.WriteFile(*catalogPath, out, 0o644); err != nil {
					return fmt.Errorf("write catalog %s: %w", *catalogPath, err)
				}
				fmt.Printf("✓ %s (добавлено %d задач)\n", *catalogPath, added)
			}
		} else {
			fmt.Println("ℹ️  Нет новых задач для импорта")
		}
	}

	// Диагностика.
	printDiagnostics(cat, progress, warns)

	if *check {
		return nil
	}

	// Рендер и запись.
	if err := os.MkdirAll(*outDir, 0o755); err != nil {
		return fmt.Errorf("create out-dir %s: %w", *outDir, err)
	}
	for i := range cat.Blocks {
		b := &cat.Blocks[i]
		content := RenderPlan(b, progress)
		outPath := filepath.Join(*outDir, b.PlanFile)
		if *dryRun {
			old, _ := os.ReadFile(outPath)
			if !bytes.Equal(old, []byte(content)) {
				fmt.Printf("[dry-run] %s будет перезаписан (%d байт → %d байт)\n", outPath, len(old), len(content))
			} else {
				fmt.Printf("[dry-run] %s без изменений\n", outPath)
			}
			continue
		}
		if err := os.WriteFile(outPath, []byte(content), 0o644); err != nil {
			return fmt.Errorf("write %s: %w", outPath, err)
		}
		fmt.Printf("✓ %s\n", outPath)
	}
	return nil
}

// printDiagnostics выводит сводку и предупреждения.
func printDiagnostics(cat *Catalog, progress map[int]*NoteProgress, warns []string) {
	fmt.Println("=== Диагностика ===")
	for _, w := range warns {
		fmt.Println(w)
	}

	// Orphan-заметки (есть прогресс, нет в каталоге).
	catalogIDs := make(map[int]bool)
	for _, b := range cat.Blocks {
		for _, p := range b.Problems {
			catalogIDs[p.LeetcodeID] = true
		}
	}
	var orphans []int
	for id := range progress {
		if !catalogIDs[id] {
			orphans = append(orphans, id)
		}
	}
	sort.Ints(orphans)
	if len(orphans) > 0 {
		fmt.Printf("ℹ️  Orphan-заметки (решены, нет в каталоге): %d\n", len(orphans))
		for _, id := range orphans {
			fmt.Printf("     %d — %s\n", id, progress[id].Title)
		}
	}

	// Сводка по блокам.
	fmt.Println("\n=== Сводка по блокам ===")
	for _, b := range cat.Blocks {
		solved, needRepeat := 0, 0
		for _, p := range b.Problems {
			if np, ok := progress[p.LeetcodeID]; ok {
				solved++
				if np.NeedRepeat {
					needRepeat++
				}
			}
		}
		fmt.Printf("  %-20s %d/%d решено, %d needRepeat\n", b.ID, solved, len(b.Problems), needRepeat)
	}
}

// -------------------- Импорт (--init) --------------------

// runInit парсит существующие NEETCODE_PLAN_*.md и собирает catalog.yaml.
func runInit(plansDir, catalogPath string) error {
	importer := newPlanImporter()
	if err := importer.importDir(plansDir); err != nil {
		return err
	}
	cat := importer.catalog

	out, err := yaml.Marshal(cat)
	if err != nil {
		return fmt.Errorf("marshal catalog: %w", err)
	}
	if err := os.WriteFile(catalogPath, out, 0o644); err != nil {
		return fmt.Errorf("write catalog: %w", err)
	}
	fmt.Printf("✓ catalog.yaml создан (%d блоков, %d задач)\n", len(cat.Blocks), importer.totalProblems)

	// Шпаргалка по retention.
	if len(importer.retentionHints) > 0 {
		fmt.Println("\n=== Шпаргалка: текущие retention из планов ===")
		fmt.Println("Проставь эти значения в поле retention: заметок Obsidian:")
		var ids []int
		for id := range importer.retentionHints {
			ids = append(ids, id)
		}
		sort.Ints(ids)
		for _, id := range ids {
			fmt.Printf("  %d → L%d\n", id, importer.retentionHints[id])
		}
	}
	return nil
}

// nowStamp возвращает текущую дату в формате YAML-поля (не используется напрямую,
// но держится как референс формата для будущего поля ревью).
func nowStamp() string {
	return time.Now().Format("2006-01-02")
}
