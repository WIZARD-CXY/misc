import sys

text = sys.stdin.read()
words = text.split()
wordcount = len(words)
print "the word number is " , wordcount
