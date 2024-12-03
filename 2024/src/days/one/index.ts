/**
 * ©️ 2024-Dec-03 @aaanh | All rights reserved
 * Advent of Code 2024 | Day 1
 */

import fs from "fs/promises";

export default async function One() {
  try {
    const input = await fs.readFile("src/days/one/input-2.txt", "utf-8");

    // process the lines
    const lines = input
      .split("\n")
      .map((line) => line.trimEnd().split(/\s+/))
      .map((line) => [parseInt(line[0]), parseInt(line[1])]);

    // column to row transposition
    const transposed = lines[0]
      .map((_, colIdx) => lines.map((row) => row[colIdx]))
      .map((arr) => arr.sort());

    // calc the dx of each pair
    const dx = transposed[0].map((element, idx) =>
      // console.log(element, idx, transposed[0][idx])
      Math.abs(element - transposed[1][idx])
    );

    const sum1 = dx.reduce((acc, x) => acc + x, 0);
    console.log("part 1:", sum1);

    console.log("part 2:");
    let sum2 = 0;

    // get unique values into a set then reduce the set into an key-value object
    const unique = [...new Set(transposed[0])].reduce(
      (obj: { [key: string]: number }, key) => {
        obj[String(key)] = 0;
        return obj;
      },
      {}
    );

    // count occurrences in right column (as originally formatted by the problem)
    transposed[1].forEach((value) => {
      if (unique.hasOwnProperty(String(value))) {
        unique[String(value)]++;
      }
    });

    // summing the (value * occurrence)
    sum2 = Object.entries(unique).reduce(
      (sum, [key, count]) => sum + parseInt(key) * count,
      0
    );
    console.log(sum2);
  } catch (err) {
    throw err;
  }
}
