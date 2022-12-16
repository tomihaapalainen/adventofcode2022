import ast
import functools


def compare(l, r):
    if type(l) == type(r) == int:
        return l - r
    elif type(l) == type(r) == list:
        for ll, rr in zip(l, r):
            result = compare(ll, rr)
            if result != 0:
                return result
        return len(l) - len(r)
    elif type(l) == int:
        return compare([l], r)
    else:
        return compare(l, [r])


def adv13a():
    count = 0
    index = 1

    with open("input13.txt", "r") as f:
        inputs = []
        for line in f.readlines():
            if line.strip():
                inputs.append(line)
            else:
                inputs = []
                index += 1

            if len(inputs) == 2:
                left, right = inputs
                left = ast.literal_eval(left)
                right = ast.literal_eval(right)

                result = compare(left, right)
                if result < 0:
                    count += index

    print(count)


def adv13b():
    data = [[[2]], [[6]]]
    with open("input13.txt", "r") as f:
        inputs = []
        for line in f.readlines():
            if line.strip():
                inputs.append(line)
            else:
                inputs = []

            if len(inputs) == 2:
                left, right = inputs
                left = ast.literal_eval(left)
                right = ast.literal_eval(right)
                data.append(left)
                data.append(right)

    data.sort(key=functools.cmp_to_key(compare))
    i1 = data.index([[2]]) + 1
    i2 = data.index([[6]]) + 1

    print(i1 * i2)


if __name__ == "__main__":
    adv13b()
