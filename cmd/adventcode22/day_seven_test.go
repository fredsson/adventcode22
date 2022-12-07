package main

import "testing"

func Test_AddDirentToDirectory_ShouldUpdateTotalSize(t *testing.T) {
	root := CreateDirectory("dir /", nil)

	AddDirentToDirectory("73403 t.txt", root)

	if root.TotalSize != 73403 {
		t.Errorf("Total size of directory is incorrect expected %d but got %d", 73403, root.TotalSize)
	}

}

func Test_HandleCdCommand_ShouldNotChangeDirectoryWhenMovingToSame(t *testing.T) {
	root := CreateDirectory("dir /", nil)

	current := HandleCdCommand("$ cd /", root)

	if current != root {
		t.Error("Should still be in root directory")
	}
}

func Test_HandleCdCommand_ShouldPrioritizeMovingToChildWhenNamedSame(t *testing.T) {
	root := CreateDirectory("dir /", nil)
	expectedDirectory := CreateDirectory("dir /", root)
	root.Content = append(root.Content, expectedDirectory)

	current := HandleCdCommand("$ cd /", root)

	if current != expectedDirectory {
		t.Error("Should be in child directory")
	}
}

func Test_HandleCdCommand_ShouldMoveToCorrectSubDirectory(t *testing.T) {
	root := CreateDirectory("dir /", nil)
	expectedDirectory := CreateDirectory("dir test", root)
	root.Content = append(root.Content, expectedDirectory)

	current := HandleCdCommand("$ cd test", root)

	if current != expectedDirectory {
		t.Error("Should be in test directory")
	}
}

func Test_HandleCdCommand_ShouldMoveUpOneDirectory(t *testing.T) {
	root := CreateDirectory("dir /", nil)
	currentDirectory := CreateDirectory("dir test", root)
	root.Content = append(root.Content, currentDirectory)

	current := HandleCdCommand("$ cd ..", currentDirectory)

	if current != root {
		t.Error("Should be in root directory")
	}
}

func Test_CalculateSizeForAllDirectories_ShouldReturnRootSize(t *testing.T) {
	root := CreateDirectory("dir /", nil)
	AddDirentToDirectory("73403 t.txt", root)

	sizes := CalculateSizeForAllDirectories(root)

	if sizes["/"] != 73403 {
		t.Errorf("Total size of directory is incorrect expected %d but got %d", 73403, sizes["/"])
	}
}

func Test_CalculateSizeForAllDirectories_ShouldReturnSubDirectorySize(t *testing.T) {
	root := CreateDirectory("dir /", nil)
	subDirectory := CreateDirectory("dir test", root)
	root.Content = append(root.Content, subDirectory)
	AddDirentToDirectory("800 t.txt", root)
	AddDirentToDirectory("200 t.txt", subDirectory)

	sizes := CalculateSizeForAllDirectories(root)

	if sizes["/test"] != 200 {
		t.Errorf("Total size of directory is incorrect expected %d but got %d", 200, sizes["/test"])
	}

	if sizes["/"] != 1000 {
		t.Errorf("Total size of directory is incorrect expected %d but got %d", 1000, sizes["/"])
	}
}

func Test_CalculateSizeForAllDirectories_ShouldPropagateSizeForMultipleLevels(t *testing.T) {
	root := CreateDirectory("dir /", nil)
	subDirectory := CreateDirectory("dir test", root)
	subSubDirectory := CreateDirectory("dir another", subDirectory)
	root.Content = append(root.Content, subDirectory)
	subDirectory.Content = append(subDirectory.Content, subSubDirectory)
	AddDirentToDirectory("800 t.txt", root)
	AddDirentToDirectory("200 t.txt", subSubDirectory)

	sizes := CalculateSizeForAllDirectories(root)

	if sizes["test/another"] != 200 {
		t.Errorf("Total size of directory is incorrect expected %d but got %d", 200, sizes["test/another"])
	}

	if sizes["/test"] != 200 {
		t.Errorf("Total size of directory is incorrect expected %d but got %d", 200, sizes["/test"])
	}

	if sizes["/"] != 1000 {
		t.Errorf("Total size of directory is incorrect expected %d but got %d", 1000, sizes["/"])
	}
}
