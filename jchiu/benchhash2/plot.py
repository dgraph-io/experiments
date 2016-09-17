import matplotlib.pyplot as plt
import numpy as np

with open('results.txt') as f:
	a = f.read().splitlines()

a = [s.split() for s in a]
a = {s[0].split('-')[0]: s for s in a}

n = 20

for mapType in ('GoMap', 'GotomicMap', 'ShardedGoMap16', 'ShardedGoMap64'):
	results = []
	for i in range(n + 1):
		key = 'BenchmarkReadWrite/%d/%s' % (i, mapType)
		key2 = 'BenchmarkReadWriteControl/%d' % i
		assert key in a
		assert key2 in a
		v = float(a[key][2])
		v2 = float(a[key2][2])
		results.append(v - v2)
		
	x = np.linspace(0, 1, n + 1)
	plt.plot(x, results, label=mapType)
	plt.ylabel('ns/op')
	plt.xlabel('readFrac')

plt.legend()
plt.show()