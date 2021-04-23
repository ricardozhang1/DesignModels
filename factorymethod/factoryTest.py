# encoding: utf-8

class Base(object):
    def __init__(self):
        self.a = 0
        self.b = 0

    def set_a_value(self, v):
        self.a = v

    def set_b_value(self, v):
        self.b = v


class Plus(Base):
    def result(self):
        return self.a + self.b


class Minus(Base):
    def result(self):
        return self.a - self.b


def compute(factory, a, b):
    factory.set_a_value(a)
    factory.set_b_value(b)
    return factory.result()


if __name__ == '__main__':
    # f = Base()

    f = Plus()
    r = compute(f, 1, 2)
    print(r)

    f = Minus()
    r = compute(f, 6, 2)
    print(r)


