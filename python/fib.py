class Fibs(object):
	def __init__(self):
		self.a=0
		self.b=1
	def next(self):
		self.a ,self.b = self.b,self.a+self.b
		return self.a
	def __iter__(self):
		return self


class TestIterator:
	value = 0
	def next (self):
		self.value +=1
		if self.value > 100 : raise StopIteration
		return self.value
	def __iter__(self):
		return self

def flatten(nested):
  try:
    for sublist in nested:
      for element in flatten(sublist):
        yield element
  except TypeError:
    yield nested
if __name__ == '__main__':
	fibs = Fibs()

	for f in fibs:
		if f <10000:
			temp = f
		if f >10000:
                    print temp
                    break

	ti = TestIterator()
	print list(ti)
        print list(flatten([[[1],2],3,4,[5,[6,7]],8]))
