package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// -------------------- Импортёр существующих планов в catalog.yaml --------------------
//
// Парсит текущие NEETCODE_PLAN_*.md и собирает статичные метаданные задач.
// Соглашения о формате (выведены из реальных планов в репозитории):
//   - Строки-задачи начинаются с "| [x]" или "| [ ]".
//   - Колонка "Задача" имеет вид "**NNN. Title**" (bold) или "NNN. Title".
//   - Секция определяется заголовком "## ..." или отсутствует (→ main).
//   - Блок определяется именем файла: NEETCODE_PLAN_<n>.md.

// planImporter собирает catalog.yaml из markdown-планов.
type planImporter struct {
	catalog         Catalog
	totalProblems   int
	retentionHints  map[int]int // leetcode_id → retention (из старых меток L1-L5)
	blockByPlanFile map[string]int
}

func newPlanImporter() *planImporter {
	return &planImporter{
		retentionHints:  make(map[int]int),
		blockByPlanFile: make(map[string]int),
	}
}

// importDir обходит plansDir и парсит все NEETCODE_PLAN_*.md.
func (im *planImporter) importDir(plansDir string) error {
	entries, err := os.ReadDir(plansDir)
	if err != nil {
		return fmt.Errorf("read plans dir: %w", err)
	}
	var planFiles []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if !strings.HasPrefix(name, "NEETCODE_PLAN_") || !strings.HasSuffix(name, ".md") {
			continue
		}
		planFiles = append(planFiles, name)
	}
	if len(planFiles) == 0 {
		return fmt.Errorf("в %s нет файлов NEETCODE_PLAN_*.md", plansDir)
	}
	sort.Strings(planFiles)

	// Стабильный маппинг plan_file → block id/folder/title по известным планам.
	blockMeta := map[string]struct{ id, folder, title string }{
		"NEETCODE_PLAN_1.md": {"arrays_hashing", "Arrays_Hashing", "Arrays & Hashing"},
		"NEETCODE_PLAN_2.md": {"two_pointers", "Two_pointers", "Two Pointers"},
		"NEETCODE_PLAN_3.md": {"sliding_window", "Sliding window", "Sliding Window"},
		"NEETCODE_PLAN_4.md": {"linked_list", "Linked List", "Linked List"},
	}

	for _, name := range planFiles {
		path := filepath.Join(plansDir, name)
		block, err := im.parsePlanFile(path, name, blockMeta)
		if err != nil {
			return fmt.Errorf("parse %s: %w", name, err)
		}
		im.catalog.Blocks = append(im.catalog.Blocks, block)
		im.blockByPlanFile[name] = len(im.catalog.Blocks) - 1
	}

	im.catalog.Version = 1
	im.catalog.VaultRoot = "/home/bryack/Documents/Obsidian/bryack/014_Go/algorithms"
	return nil
}

// parsePlanFile парсит один markdown-план в Block.
func (im *planImporter) parsePlanFile(path, planFile string, meta map[string]struct{ id, folder, title string }) (Block, error) {
	f, err := os.Open(path)
	if err != nil {
		return Block{}, err
	}
	defer f.Close()

	var b Block
	b.PlanFile = planFile
	if m, ok := meta[planFile]; ok {
		b.ID = m.id
		b.Folder = m.folder
		b.Title = m.title
	} else {
		// Фолбэк: выводим id из имени файла.
		base := strings.TrimSuffix(strings.TrimPrefix(planFile, "NEETCODE_PLAN_"), ".md")
		b.ID = "plan_" + base
		b.Folder = "Plan_" + base
		b.Title = "Plan " + base
	}

	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)

	currentGroup := "main"
	currentTitle := defaultGroupTitles["main"]
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		trim := strings.TrimSpace(line)

		// Заголовок секции.
		if strings.HasPrefix(trim, "## ") {
			heading := strings.TrimSpace(strings.TrimPrefix(trim, "## "))
			if g := guessGroup(heading); g != "" {
				currentGroup = g
				currentTitle = heading
			} else {
				// Неизвестный заголовок секции — считаем частью main.
				currentGroup = "main"
				currentTitle = heading
			}
			continue
		}

		// Строка таблицы с задачей.
		if !strings.HasPrefix(trim, "|") {
			continue
		}
		if isTableSeparator(trim) {
			continue
		}
		p, retention, ok := parsePlanRow(trim)
		if !ok {
			continue
		}
		// Дедупликация: одна задача может повторяться в разных секциях плана
		// (например, "Палиндромы" содержит повтор 680 из основного блока).
		// В каталоге оставляем только первое вхождение — оно определяет группу.
		if problemExists(b.Problems, p.LeetcodeID) {
			continue
		}
		p.Group = currentGroup
		// Секции в каталоге (для рендера заголовков).
		if !sectionExists(b.Sections, currentGroup) {
			b.Sections = append(b.Sections, Section{Group: currentGroup, Title: currentTitle})
		}
		b.Problems = append(b.Problems, p)
		im.totalProblems++
		if retention > 0 {
			im.retentionHints[p.LeetcodeID] = retention
		}
	}
	if err := scanner.Err(); err != nil {
		return Block{}, err
	}
	return b, nil
}

// isTableSeparator определяет строку-разделитель "| :---: | ... |".
func isTableSeparator(line string) bool {
	if !strings.HasPrefix(line, "|") {
		return false
	}
	s := strings.ReplaceAll(line, " ", "")
	s = strings.ReplaceAll(s, "|", "")
	return strings.Trim(s, "-:") == ""
}

// parsePlanRow разбирает строку таблицы в Problem + retention (если есть).
// Формат колонок в текущих планах (Plan_2/3/4): Статус | Задача | Сложность | Тема | Дата | Retention | Ревью
// В Plan_1 колонок меньше: Статус | Задача | Сложность | Тема | Дата | Ревью.
func parsePlanRow(line string) (Problem, int, bool) {
	cells := splitTableRow(line)
	if len(cells) < 4 {
		return Problem{}, 0, false
	}

	// Колонка 0: статус "[x]"/"[ ]".
	status := strings.TrimSpace(cells[0])
	if status != "[x]" && status != "[ ]" {
		// Возможно, это заголовок таблицы или легенды.
		return Problem{}, 0, false
	}

	// Колонка 1: задача "**NNN. Title**" или "NNN. Title".
	taskCell := strings.TrimSpace(cells[1])
	taskCell = strings.Trim(taskCell, "*")
	id, title, ok := splitTaskCell(taskCell)
	if !ok {
		return Problem{}, 0, false
	}

	// Колонка 2: сложность (Easy/Medium/Hard) — у Plan_1 может отличаться.
	diff := normalizeDifficulty(strings.TrimSpace(cells[2]))

	// Колонка 3: паттерн.
	pattern := strings.TrimSpace(cells[3])
	pattern = strings.Trim(pattern, "*")

	p := Problem{
		LeetcodeID: id,
		Title:      title,
		Difficulty: diff,
		Pattern:    pattern,
	}

	// Retention ищем в оставшихся колонках по виду "L4", "🟢 L4" и т.п.
	retention := 0
	for _, c := range cells[4:] {
		if r := extractRetention(strings.TrimSpace(c)); r > 0 {
			retention = r
			break
		}
	}
	return p, retention, true
}

// splitTableRow делит строку "| a | b | c |" на ["a","b","c"].
func splitTableRow(line string) []string {
	s := strings.TrimSpace(line)
	s = strings.TrimPrefix(s, "|")
	s = strings.TrimSuffix(s, "|")
	parts := strings.Split(s, "|")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

// taskCellRe — разбирает "NNN. Title" (опционально с префиксом "Practice:").
var taskCellRe = regexp.MustCompile(`^(?:\*{0,2})?(?:Practice:\s*)?(\d+)\.\s*(.+?)(?:\*{0,2})?$`)

// splitTaskCell извлекает leetcode_id и название из ячейки задачи.
func splitTaskCell(cell string) (int, string, bool) {
	cell = strings.TrimSpace(cell)
	cell = strings.Trim(cell, "*")
	cell = strings.TrimSpace(cell)
	// Убираем возможный префикс "Practice: ".
	cell = strings.TrimPrefix(cell, "Practice: ")
	m := taskCellRe.FindStringSubmatch(cell)
	if m == nil {
		return 0, "", false
	}
	id, err := strconv.Atoi(m[1])
	if err != nil {
		return 0, "", false
	}
	title := strings.TrimSpace(m[2])
	title = strings.Trim(title, "*")
	return id, title, true
}

// normalizeDifficulty приводит сложность к каноническому виду.
func normalizeDifficulty(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "*")
	switch strings.ToLower(s) {
	case "easy":
		return "Easy"
	case "medium":
		return "Medium"
	case "hard":
		return "Hard"
	default:
		return s
	}
}

// retentionRe — извлекает число из метки вида "L4", "🟢 L4", "⚪ L5".
var retentionRe = regexp.MustCompile(`L([1-5])`)

// extractRetention возвращает 1-5 из строки с retention-меткой или 0.
func extractRetention(s string) int {
	if s == "" {
		return 0
	}
	m := retentionRe.FindStringSubmatch(s)
	if m == nil {
		return 0
	}
	r, err := strconv.Atoi(m[1])
	if err != nil {
		return 0
	}
	return r
}

// guessGroup определяет группу по заголовку секции плана.
func guessGroup(heading string) string {
	h := strings.ToLower(heading)
	switch {
	case strings.Contains(h, "bonus"):
		return "bonus"
	case strings.Contains(h, "яндекс") || strings.Contains(h, "yandex"):
		return "yandex"
	case strings.Contains(h, "основной"), strings.Contains(h, "главный"), strings.Contains(h, "main"):
		return "main"
	case strings.Contains(h, "дополнительн"), strings.Contains(h, "стек"), strings.Contains(h, "очеред"), strings.Contains(h, "практик"):
		// "Дополнительная практика: Палиндромы" и т.п. — относим к yandex-группе
		// (в текущих планах такие секции содержат яндексовские/практические задачи).
		return "yandex"
	default:
		// Если заголовок содержит название блока — основной.
		return "main"
	}
}

// sectionExists проверяет, есть ли уже группа в секциях блока.
func sectionExists(sections []Section, group string) bool {
	for _, s := range sections {
		if s.Group == group {
			return true
		}
	}
	return false
}

// problemExists проверяет, есть ли уже задача с данным leetcode_id в блоке.
func problemExists(problems []Problem, id int) bool {
	for _, p := range problems {
		if p.LeetcodeID == id {
			return true
		}
	}
	return false
}
