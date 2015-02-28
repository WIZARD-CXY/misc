class Bird:
    def __init__(self):
        self.hungry = True
    def eat(self):
        if self.hungry:
            print "AAAA"
            self.hungry = False
        else:
            print "No,thanks"
b = Bird()
b.eat()
