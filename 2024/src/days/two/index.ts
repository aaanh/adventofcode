import fs from "fs";

export default function Two() {
  const input = fs.readFileSync("src/days/two/input.txt", "utf-8");
  const matrix = input
    .split("\n")
    .map((line) => line.split(" ").map((element) => parseInt(element)));

  console.log("lines: ", matrix.length);

  let safe = matrix.length;

  for (let rowIdx = 0; rowIdx < matrix.length; rowIdx++) {
    console.log();
    let isIncreasing = false;
    for (let colIdx = 0; colIdx < matrix[rowIdx].length - 1; colIdx++) {
      let dx = matrix[rowIdx][colIdx] - matrix[rowIdx][colIdx + 1];
      let abs = Math.abs(dx);

      process.stdout.write(
        matrix[rowIdx][colIdx].toString() +
          "-" +
          matrix[rowIdx][colIdx + 1].toString() +
          "=" +
          dx.toString() +
          " "
      );

      if (abs > 3 || abs == 0) {
        safe--;
        console.log("broke");
        break;
      }

      if (dx < 0) {
        isIncreasing = true;
      }

      if (dx > 0 && isIncreasing) {
        safe--;
        console.log("broke");
        break;
      }
    }
  }

  console.log(safe);
}
