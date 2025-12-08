import itertools as it
import math
import time

start = time.time()

circuits = {b: {b} for b in map(eval, open("input.txt"))}

pairs = sorted(it.combinations(circuits, 2), key=lambda x: math.dist(*x))

for i, (box1, box2) in enumerate(pairs):
    for c in circuits:
        if box1 in circuits[c]:
            cir1 = c
        if box2 in circuits[c]:
            cir2 = c

    if cir1 != cir2:
        circuits[cir1] |= circuits[cir2]
        del circuits[cir2]

    if i + 1 == 1000:
        n = sorted(len(circuits[b]) for b in circuits)
        print(n[-3] * n[-2] * n[-1])

    if len(circuits) == 1:
        print(box1[0] * box2[0])
        break

end = time.time()
print(end - start)
