package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

// testdataDir — абсолютный путь к testdata.
func testdataDir(t *testing.T) string {
	t.Helper()
	abs, err := filepath.Abs("testdata")
	require.NoError(t, err)
	return abs
}

// -------------------- LoadCatalog / Validate --------------------

func TestLoadCatalog_Valid(t *testing.T) {
	cat, err := LoadCatalog(filepath.Join(testdataDir(t), "catalog.yaml"))
	require.NoError(t, err)
	assert.Equal(t, 2, len(cat.Blocks))
	assert.Equal(t, "arrays_hashing", cat.Blocks[0].ID)
	assert.Equal(t, 3, len(cat.Blocks[0].Problems))
}

func TestValidate_DuplicateWithinBlock(t *testing.T) {
	cat := &Catalog{
		VaultRoot: "/tmp",
		Blocks: []Block{{
			ID:       "b1",
			Folder:   "F1",
			PlanFile: "P1.md",
			Problems: []Problem{
				{LeetcodeID: 1, Title: "A", Difficulty: "Easy"},
				{LeetcodeID: 1, Title: "A again", Difficulty: "Easy"}, // дубль внутри блока
			},
		}},
	}
	err := cat.Validate()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "duplicated within block")
}

func TestValidate_DuplicateAcrossBlocks_Allowed(t *testing.T) {
	// 350 в двух блоках — разрешено (задача изучается разными паттернами).
	cat := &Catalog{
		VaultRoot: "/tmp",
		Blocks: []Block{
			{
				ID: "arrays_hashing", Folder: "Arrays_Hashing", PlanFile: "P1.md",
				Problems: []Problem{{LeetcodeID: 350, Title: "Intersection II", Difficulty: "Easy"}},
			},
			{
				ID: "two_pointers", Folder: "Two_pointers", PlanFile: "P2.md",
				Problems: []Problem{{LeetcodeID: 350, Title: "Intersection II", Difficulty: "Easy"}},
			},
		},
	}
	assert.NoError(t, cat.Validate())
}

func TestValidate_DuplicateFolder(t *testing.T) {
	cat := &Catalog{
		VaultRoot: "/tmp",
		Blocks: []Block{
			{ID: "b1", Folder: "Same", PlanFile: "P1.md"},
			{ID: "b2", Folder: "Same", PlanFile: "P2.md"},
		},
	}
	err := cat.Validate()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "folder")
}

func TestValidate_InvalidDifficulty(t *testing.T) {
	cat := &Catalog{
		VaultRoot: "/tmp",
		Blocks: []Block{{
			ID: "b1", Folder: "F1", PlanFile: "P1.md",
			Problems: []Problem{{LeetcodeID: 1, Difficulty: "Trivial"}},
		}},
	}
	err := cat.Validate()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid difficulty")
}

// -------------------- ScanVault --------------------

func TestScanVault_ParsesProgress(t *testing.T) {
	vault := filepath.Join(testdataDir(t), "vault")
	progress, warns, err := ScanVault(vault)
	require.NoError(t, err)

	// 217 — решена, retention=4, две даты.
	np, ok := progress[217]
	require.True(t, ok, "217 должна быть распознана")
	assert.Equal(t, 4, int(np.Retention))
	assert.False(t, np.NeedRepeat)
	assert.Equal(t, []string{"2026-04-24", "2026-06-28"}, np.SolvedDates)
	assert.Equal(t, "Arrays_Hashing", np.Folder)

	// 1 — needRepeat=true, retention не проставлен (=0).
	np1, ok := progress[1]
	require.True(t, ok)
	assert.True(t, np1.NeedRepeat)
	assert.Equal(t, 0, int(np1.Retention))

	// 344 — retention=5.
	np344, ok := progress[344]
	require.True(t, ok)
	assert.Equal(t, 5, int(np344.Retention))
	assert.Equal(t, "Two_pointers", np344.Folder)

	// Orphan 999 присутствует.
	_, ok = progress[999]
	assert.True(t, ok, "orphan 999 должна быть распознана")

	// Предупреждения: хотя бы одна заметка без номера в имени.
	assert.NotEmpty(t, warns)
	joined := strings.Join(warns, "\n")
	assert.Contains(t, joined, "без leetcode_id в имени")
}

// -------------------- extractFrontmatter --------------------

func TestExtractFrontmatter(t *testing.T) {
	lines := []string{
		"---",
		"Type:",
		"  - exercise",
		"created: 2026-07-10",
		"---",
		"# body",
	}
	fm, ok := extractFrontmatter(lines)
	require.True(t, ok)
	assert.Contains(t, fm, "exercise")
	assert.Contains(t, fm, "2026-07-10")
	assert.NotContains(t, fm, "# body")
}

func TestExtractFrontmatter_NoFrontmatter(t *testing.T) {
	lines := []string{"# Заголовок", "Текст"}
	_, ok := extractFrontmatter(lines)
	assert.False(t, ok)
}

func TestExtractFrontmatter_Unclosed(t *testing.T) {
	lines := []string{"---", "created: 2026-07-10"}
	_, ok := extractFrontmatter(lines)
	assert.False(t, ok, "незакрытый frontmatter должен давать false")
}

// -------------------- isExerciseNote --------------------

func TestIsExerciseNote(t *testing.T) {
	assert.True(t, isExerciseNote("Type:\n  - exercise\ntags:\n  - algorithms"))
	assert.False(t, isExerciseNote("Type:\n  - note\ntags:\n  - theory"))
	assert.False(t, isExerciseNote("created: 2026-07-10"))
}

// -------------------- retentionEmoji / firstSolvedDate --------------------

func TestRetentionEmoji(t *testing.T) {
	cases := map[flexInt]string{
		5: "⚪ L5",
		4: "🟢 L4",
		3: "🟡 L3",
		2: "🟠 L2",
		1: "🔴 L1",
		0: "",
		9: "",
	}
	for in, want := range cases {
		assert.Equal(t, want, retentionEmoji(in), "retention=%d", in)
	}
}

// TestFlexInt_UnmarshalYAML проверяет, что retention корректно парсится
// из разных форматов (int, строка с кавычками, строка с префиксом "L").
func TestFlexInt_UnmarshalYAML(t *testing.T) {
	cases := map[string]flexInt{
		`retention: 3`:         3,
		`retention: "3"`:       3,
		`retention: "L4"`:      4,
		`retention: "level 2"`: 2,
		`retention: ""`:        0,
	}
	for yamlStr, want := range cases {
		var np NoteProgress
		err := yaml.Unmarshal([]byte(yamlStr), &np)
		require.NoError(t, err, "input=%q", yamlStr)
		assert.Equal(t, int(want), int(np.Retention), "input=%q", yamlStr)
	}
}

func TestFirstSolvedDate(t *testing.T) {
	np := &NoteProgress{SolvedDates: []string{"2026-06-28", "2026-04-24", "2026-05-01"}}
	assert.Equal(t, "2026-04-24", firstSolvedDate(np))

	assert.Equal(t, "", firstSolvedDate(nil))
	assert.Equal(t, "", firstSolvedDate(&NoteProgress{}))
}

// -------------------- RenderPlan --------------------

func TestRenderPlan_OrphanSection(t *testing.T) {
	cat, err := LoadCatalog(filepath.Join(testdataDir(t), "catalog.yaml"))
	require.NoError(t, err)
	vault := filepath.Join(testdataDir(t), "vault")
	progress, _, err := ScanVault(vault)
	require.NoError(t, err)

	// Блок two_pointers содержит 999 как orphan.
	var tp *Block
	for i := range cat.Blocks {
		if cat.Blocks[i].ID == "two_pointers" {
			tp = &cat.Blocks[i]
		}
	}
	require.NotNil(t, tp)

	out := RenderPlan(tp, progress)
	// Orphan-секция должна присутствовать.
	assert.Contains(t, out, "Вне каталога (Extra)")
	assert.Contains(t, out, "999. Orphan Task")
	// Задача 350 из каталога тоже рендерится.
	assert.Contains(t, out, "350. Intersection")
	// Подсказка про catalog.yaml.
	assert.Contains(t, out, "catalog.yaml")
}

func TestRenderPlan_StatusAndRetention(t *testing.T) {
	cat, err := LoadCatalog(filepath.Join(testdataDir(t), "catalog.yaml"))
	require.NoError(t, err)
	vault := filepath.Join(testdataDir(t), "vault")
	progress, _, err := ScanVault(vault)
	require.NoError(t, err)

	var ah *Block
	for i := range cat.Blocks {
		if cat.Blocks[i].ID == "arrays_hashing" {
			ah = &cat.Blocks[i]
		}
	}
	require.NotNil(t, ah)

	out := RenderPlan(ah, progress)
	// 217 — решена ([x]), retention L4.
	assert.Contains(t, out, "🟢 L4")
	// Статус решённой задачи.
	assert.Contains(t, out, "[x]")
	// Группы: основной блок и Яндекс.
	assert.Contains(t, out, "Основной блок")
	assert.Contains(t, out, "Задачи с собеседований Яндекс")
	// Two Sum помечена needRepeat: true.
	assert.Contains(t, out, "true")
	// Колонка количества решений присутствует.
	assert.Contains(t, out, "Решений")
}

// -------------------- parsePlanRow / parsePlanFile (для --init) --------------------

func TestParsePlanRow_Standard(t *testing.T) {
	// Строка в формате Plan_2/3/4.
	row := "| [x] | **680. Valid Palindrome II** | Easy | Two Pointers (left/right, skip check) | 2026-06-01 | 🟢 L4 | |"
	p, ret, ok := parsePlanRow(row)
	require.True(t, ok)
	assert.Equal(t, 680, p.LeetcodeID)
	assert.Equal(t, "Valid Palindrome II", p.Title)
	assert.Equal(t, "Easy", p.Difficulty)
	assert.Equal(t, "Two Pointers (left/right, skip check)", p.Pattern)
	assert.Equal(t, 4, ret)
}

func TestParsePlanRow_Practice(t *testing.T) {
	row := "| [x] | **Practice: 350. Intersection of Two Arrays II** | Easy | Frequency Counting / Intersection | 2026-04-25 | |"
	p, _, ok := parsePlanRow(row)
	require.True(t, ok)
	assert.Equal(t, 350, p.LeetcodeID)
	assert.Equal(t, "Intersection of Two Arrays II", p.Title)
}

func TestParsePlanRow_NotATaskRow(t *testing.T) {
	// Строка легенды/заголовка — не задача.
	_, _, ok := parsePlanRow("| 🟢 **L4** | Решила сама, быстро, чистый код | Без подсказок |")
	assert.False(t, ok)
}

func TestParsePlanRow_TableSeparator(t *testing.T) {
	assert.True(t, isTableSeparator("| :---: | :--- | :---: | :--- |"))
	assert.False(t, isTableSeparator("| [x] | **1. Two Sum** | Easy | HashMap |"))
}

func TestGuessGroup(t *testing.T) {
	cases := map[string]string{
		"Основной блок":                            "main",
		"Bonus (NeetCode 150, уже решены)":         "bonus",
		"Задачи с собеседований Яндекс":            "yandex",
		"Дополнительная практика: Палиндромы":      "yandex",
		"Дополнительная практика: Стеки и Очереди": "yandex",
	}
	for heading, want := range cases {
		assert.Equal(t, want, guessGroup(heading), "heading=%q", heading)
	}
}

// -------------------- Интеграционный тест runInit --------------------

func TestRunInit_EndToEnd(t *testing.T) {
	// Подготовим временную директорию с одним планом и запустим --init.
	tmp := t.TempDir()
	plan := `# NeetCode Roadmap — Arrays & Hashing

## Основной блок

| Статус | Задача | Сложность | Тема / Паттерн | Дата | Retention | Ревью |
| :---: | :--- | :---: | :--- | :--- | :--- | :--- |
| [x] | **217. Contains Duplicate** | Easy | Hash Set | 2026-04-24 | 🟢 L4 | |
| [x] | **1. Two Sum** | Easy | HashMap | 2026-04-24 | 🟢 L4 | |

## Задачи с собеседований Яндекс

| Статус | Задача | Сложность | Тема / Паттерн | Дата | Retention | Ревью |
| :---: | :--- | :---: | :--- | :--- | :--- | :--- |
| [x] | **2215. Find the Difference of Two Arrays** | Easy | HashSet Difference | 2026-07-13 | 🟠 L2 | |
`
	planPath := filepath.Join(tmp, "NEETCODE_PLAN_1.md")
	require.NoError(t, os.WriteFile(planPath, []byte(plan), 0o644))

	catalogPath := filepath.Join(tmp, "catalog.yaml")
	im := newPlanImporter()
	require.NoError(t, im.importDir(tmp))
	require.Equal(t, 1, len(im.catalog.Blocks))
	b := im.catalog.Blocks[0]
	assert.Equal(t, "arrays_hashing", b.ID)
	require.Equal(t, 3, len(b.Problems))
	// Порядок сохранён.
	assert.Equal(t, 217, b.Problems[0].LeetcodeID)
	assert.Equal(t, "main", b.Problems[0].Group)
	assert.Equal(t, 2215, b.Problems[2].LeetcodeID)
	assert.Equal(t, "yandex", b.Problems[2].Group)
	// Retention-шпаргалка собрана.
	assert.Equal(t, 4, im.retentionHints[217])
	assert.Equal(t, 2, im.retentionHints[2215])

	// Теперь запишем каталог и прочитаем обратно — должен валидироваться.
	out, err := yaml.Marshal(im.catalog)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(catalogPath, out, 0o644))
	_, err = LoadCatalog(catalogPath)
	assert.NoError(t, err)
}

// -------------------- ImportOrphans --------------------

func TestImportOrphans(t *testing.T) {
	tests := []struct {
		name         string
		initialCat   *Catalog
		progress     map[int]*NoteProgress
		wantAdded    int
		checkCatalog func(t *testing.T, cat *Catalog)
	}{
		{
			name: "imports orphan into matching block and assigns group main",
			initialCat: &Catalog{
				Blocks: []Block{
					{
						ID:       "arrays_hashing",
						Folder:   "Arrays_Hashing",
						PlanFile: "NEETCODE_PLAN_1.md",
						Problems: []Problem{
							{LeetcodeID: 1, Title: "Two Sum", Difficulty: "Easy", Group: "main"},
						},
					},
				},
			},
			progress: map[int]*NoteProgress{
				9: {
					LeetcodeID: 9,
					Title:      "Palindrome Number",
					Folder:     "Arrays_Hashing",
					Difficulty: "Easy",
					Algorithms: []string{"array", "math"},
				},
			},
			wantAdded: 1,
			checkCatalog: func(t *testing.T, cat *Catalog) {
				b := cat.Blocks[0]
				require.Equal(t, 2, len(b.Problems))
				p := b.Problems[1]
				assert.Equal(t, 9, p.LeetcodeID)
				assert.Equal(t, "Palindrome Number", p.Title)
				assert.Equal(t, "Easy", p.Difficulty)
				assert.Equal(t, "array, math", p.Pattern)
				assert.Equal(t, "main", p.Group)
			},
		},
		{
			name: "imports orphan with nested folder matching parent block folder",
			initialCat: &Catalog{
				Blocks: []Block{
					{
						ID:       "trees",
						Folder:   "Binary Tree",
						PlanFile: "NEETCODE_PLAN_7.md",
						Problems: nil,
					},
				},
			},
			progress: map[int]*NoteProgress{
				98: {
					LeetcodeID: 98,
					Title:      "Validate Binary Search Tree",
					Folder:     "Binary Tree/Binary Search Tree",
					Difficulty: "Medium",
					Algorithms: []string{"bst", "dfs"},
				},
			},
			wantAdded: 1,
			checkCatalog: func(t *testing.T, cat *Catalog) {
				b := cat.Blocks[0]
				require.Equal(t, 1, len(b.Problems))
				assert.Equal(t, 98, b.Problems[0].LeetcodeID)
				assert.Equal(t, "Medium", b.Problems[0].Difficulty)
			},
		},
		{
			name: "skips orphan when already in catalog",
			initialCat: &Catalog{
				Blocks: []Block{
					{
						ID:       "arrays_hashing",
						Folder:   "Arrays_Hashing",
						PlanFile: "NEETCODE_PLAN_1.md",
						Problems: []Problem{
							{LeetcodeID: 9, Title: "Palindrome Number", Difficulty: "Easy", Group: "main"},
						},
					},
				},
			},
			progress: map[int]*NoteProgress{
				9: {
					LeetcodeID: 9,
					Title:      "Palindrome Number",
					Folder:     "Arrays_Hashing",
					Difficulty: "Easy",
				},
			},
			wantAdded: 0,
			checkCatalog: func(t *testing.T, cat *Catalog) {
				assert.Equal(t, 1, len(cat.Blocks[0].Problems))
			},
		},
		{
			name: "skips orphan when folder does not match any block",
			initialCat: &Catalog{
				Blocks: []Block{
					{
						ID:       "arrays_hashing",
						Folder:   "Arrays_Hashing",
						PlanFile: "NEETCODE_PLAN_1.md",
					},
				},
			},
			progress: map[int]*NoteProgress{
				414: {
					LeetcodeID: 414,
					Title:      "Third Maximum Number",
					Folder:     "VK Education/tasks",
					Difficulty: "Easy",
				},
				12: {
					LeetcodeID: 12,
					Title:      "Integer to Roman",
					Folder:     "tasks",
					Difficulty: "Medium",
				},
			},
			wantAdded: 0,
			checkCatalog: func(t *testing.T, cat *Catalog) {
				assert.Equal(t, 0, len(cat.Blocks[0].Problems))
			},
		},
		{
			name: "defaults empty difficulty to Easy",
			initialCat: &Catalog{
				Blocks: []Block{
					{
						ID:       "arrays_hashing",
						Folder:   "Arrays_Hashing",
						PlanFile: "NEETCODE_PLAN_1.md",
					},
				},
			},
			progress: map[int]*NoteProgress{
				58: {
					LeetcodeID: 58,
					Title:      "Length of Last Word",
					Folder:     "Arrays_Hashing",
					Difficulty: "", // empty
				},
			},
			wantAdded: 1,
			checkCatalog: func(t *testing.T, cat *Catalog) {
				require.Equal(t, 1, len(cat.Blocks[0].Problems))
				assert.Equal(t, "Easy", cat.Blocks[0].Problems[0].Difficulty)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			added := ImportOrphans(tt.initialCat, tt.progress)
			assert.Equal(t, tt.wantAdded, added)
			if tt.checkCatalog != nil {
				tt.checkCatalog(t, tt.initialCat)
			}
		})
	}
}

func TestRun_ImportOrphans(t *testing.T) {
	tmp := t.TempDir()
	vaultDir := filepath.Join(tmp, "vault", "Arrays_Hashing")
	require.NoError(t, os.MkdirAll(vaultDir, 0o755))

	noteContent := `---
Type:
  - exercise
tags:
  - algorithms
difficulty: Easy
algorithms:
  - array
  - math
---
`
	notePath := filepath.Join(vaultDir, "9. Palindrome Number.md")
	require.NoError(t, os.WriteFile(notePath, []byte(noteContent), 0o644))

	catContent := `version: 1
vault_root: ` + filepath.Join(tmp, "vault") + `
blocks:
  - id: arrays_hashing
    title: Arrays & Hashing
    folder: Arrays_Hashing
    plan_file: NEETCODE_PLAN_1.md
    problems: []
`
	catalogPath := filepath.Join(tmp, "catalog.yaml")
	require.NoError(t, os.WriteFile(catalogPath, []byte(catContent), 0o644))

	err := run([]string{
		"-catalog", catalogPath,
		"-out-dir", tmp,
		"-import-orphans",
	})
	require.NoError(t, err)

	updatedCat, err := LoadCatalog(catalogPath)
	require.NoError(t, err)
	require.Equal(t, 1, len(updatedCat.Blocks[0].Problems))
	assert.Equal(t, 9, updatedCat.Blocks[0].Problems[0].LeetcodeID)
	assert.Equal(t, "Palindrome Number", updatedCat.Blocks[0].Problems[0].Title)

	planData, err := os.ReadFile(filepath.Join(tmp, "NEETCODE_PLAN_1.md"))
	require.NoError(t, err)
	assert.Contains(t, string(planData), "Palindrome Number")
}
