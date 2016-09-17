import sys

assert(len(sys.argv) == 2)

filename = sys.argv[1]
print 'Opening go interface file: %s' % filename

with open(filename) as f:
	lines = f.read().splitlines()
	
for i, s in enumerate(lines):
	s = s.strip()
	if s == '*/':
		i_close = i
	elif s == 'import "C"':
		i_import = i
		break

# lines[i_close] is the last '*/' before we saw 'import "C"'.
assert(i_import and i_close)

with open('cgo_flags.txt') as f:
	to_add = f.read().splitlines()
	
with open(filename, 'w') as f:
	f.write('\n'.join(lines[:i_close]))
	f.write('\n\n')
	f.write('\n'.join(to_add))
	f.write('\n\n')
	f.write('\n'.join(lines[i_close:]))


print 'Wrote to go interface file: %s' % filename