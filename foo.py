
with open('input.txt') as f:
    adapters = [int(ln.strip()) for ln in f.readlines()]
    adapters.sort()
    switch = adapters[-1]+3
    adapters.append(switch)
    paths = {0:1}
    for i in adapters:
        paths[i] = 0
        for j in (1,2,3):
            paths[i] += paths.get(i-j,0)
    print(paths[switch])
