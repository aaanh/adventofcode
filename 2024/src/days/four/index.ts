import fs from "fs";

export default function Four() {
  const input = fs.readFileSync("src/days/four/input.txt", "utf-8");
  const lines = input.split("\n");

  const matrix = lines.map((line) => [...line.split("")]);

  console.log("lines: ", lines.length);
  console.log("columns: ", matrix[0].length);

  /**
   * row | column ----
   * [0] | 128 129 130
   * [1] | 128 129 130
   * [2] | 128 129 130
   */

  const word = "XMAS";
  const wordLength = word.length;

  let matches = 0;

  const directions = [
    [-1, -1], // up-left
    [-1, 0], // up
    [-1, 1], // up-right
    [0, -1], // left
    [0, 1], // right
    [1, -1], // down-left
    [1, 0], // down
    [1, 1], // down-right
  ];

  const checkDirection = (
    startRow: number,
    startCol: number,
    dirRow: number,
    dirCol: number
  ): boolean => {
    for (let i = 0; i < wordLength; i++) {
      const currentRow = startRow + dirRow * i;
      const currentCol = startCol + dirCol * i;

      // Check bounds
      if (
        currentRow < 0 ||
        currentRow >= matrix.length ||
        currentCol < 0 ||
        currentCol >= matrix[0].length
      ) {
        return false;
      }

      // Check character match
      if (matrix[currentRow][currentCol] !== word[i]) {
        return false;
      }
    }
    return true;
  };

  for (let rowIdx = 0; rowIdx < matrix.length; rowIdx++) {
    for (let colIdx = 0; colIdx < matrix[rowIdx].length; colIdx++) {
      if (matrix[rowIdx][colIdx] === "X") {
        // Check all 8 directions from this starting point
        for (const [dirRow, dirCol] of directions) {
          if (checkDirection(rowIdx, colIdx, dirRow, dirCol)) {
            matches++;
          }
        }
      }
    }
  }

  console.log("part 1 matches: ", matches);

  let matches_2 = 0;

  /**
   * M * S
   * * A *
   * M * S
   */

  const checkXPattern = (startRow: number, startCol: number): boolean => {
    // Check bounds first
    if (
      startRow + 2 >= matrix.length ||
      startCol - 1 < 0 ||
      startCol + 1 >= matrix[0].length
    ) {
      return false;
    }

    // First check if center is 'A'
    if (matrix[startRow + 1][startCol] !== "A") {
      return false;
    }

    // Check diagonal patterns (M's on same diagonal)
    const leftDiagonalMs =
      matrix[startRow][startCol - 1] === "M" &&
      matrix[startRow + 2][startCol - 1] === "M";
    const rightDiagonalSs =
      matrix[startRow][startCol + 1] === "S" &&
      matrix[startRow + 2][startCol + 1] === "S";

    const rightDiagonalMs =
      matrix[startRow][startCol + 1] === "M" &&
      matrix[startRow + 2][startCol + 1] === "M";
    const leftDiagonalSs =
      matrix[startRow][startCol - 1] === "S" &&
      matrix[startRow + 2][startCol - 1] === "S";

    // Check horizontal patterns (M's on same row)
    const topRowMs =
      matrix[startRow][startCol - 1] === "M" &&
      matrix[startRow][startCol + 1] === "M";
    const bottomSs =
      matrix[startRow + 2][startCol - 1] === "S" &&
      matrix[startRow + 2][startCol + 1] === "S";

    const bottomRowMs =
      matrix[startRow + 2][startCol - 1] === "M" &&
      matrix[startRow + 2][startCol + 1] === "M";
    const topSs =
      matrix[startRow][startCol - 1] === "S" &&
      matrix[startRow][startCol + 1] === "S";

    return (
      (leftDiagonalMs && rightDiagonalSs) || // Diagonal pattern 1
      (rightDiagonalMs && leftDiagonalSs) || // Diagonal pattern 2
      (topRowMs && bottomSs) || // Horizontal pattern 1
      (bottomRowMs && topSs) // Horizontal pattern 2
    );
  };

  // Iterate through the matrix looking for potential X patterns
  for (let rowIdx = 0; rowIdx < matrix.length - 2; rowIdx++) {
    for (let colIdx = 1; colIdx < matrix[rowIdx].length - 1; colIdx++) {
      if (checkXPattern(rowIdx, colIdx)) {
        matches_2++;
      }
    }
  }

  console.log("part 2 matches: ", matches_2);
}
