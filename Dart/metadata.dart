
class Television {
  @Deprecated('Use turnOn instead')
  void activate() {
    turnOn();
  }

  void turnOn() {
    print('on!');
  }
}

class SmartTelevision extends Television {
  @override
  void turnOn() {
    print('Welcome to the smart TV!');
  }
}

void main() {
  var tv = SmartTelevision();
  tv.activate();
  tv.turnOn();

  var tv2 = Television();
  tv2.activate();
  tv2.turnOn();
}

// The @deprecated annotation can be used to mark a class, method, or variable as deprecated.
// The @override annotation can be used to indicate that a member is intended to override a member of its superclass.