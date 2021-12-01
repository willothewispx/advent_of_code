import unittest


class Solution:
    def __init__(self) -> None:
        pass

    def solve_part_one(self, input: list[int]) -> int:
        sol = 0
        for i in range(1, len(input)):
            if input[i-1] < input[i]:
                sol += 1
        return sol

    def solve_part_two(self, input: list[int]) -> int:
        if len(input) <= 2:
            return 0

        sol = 0
        for i in range(3, len(input)):
            if sum(input[i-3:i]) < sum(input[i-2:i+1]):
                sol += 1
        return sol

    @staticmethod
    def path_to_array(path: str) -> list[int]:
        output = []
        with open(path, "r") as input_file:
            for line in input_file:
                output.append(int(line))
        return output


class TestSolution(unittest.TestCase):

    def test_part_one(self):
        input = [199, 200, 208, 210, 200, 207, 240, 269, 260, 263]
        self.assertEqual(Solution().solve_part_one(input), 7)

    def test_part_two(self):
        input = [199, 200, 208, 210, 200, 207, 240, 269, 260, 263]
        self.assertEqual(Solution().solve_part_two(input), 5)


if __name__ == "__main__":
    part_one = Solution().solve_part_one(Solution.path_to_array("input.txt"))
    print(f"Part 1: {part_one}")
    part_two = Solution().solve_part_two(Solution.path_to_array("input.txt"))
    print(f"Part 1: {part_two}")
