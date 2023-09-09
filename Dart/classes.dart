
abstract class Vehicle {
    void steer();
    void brake();
}
class Car extends Vehicle {
    void steer() {
        print("Car is steering");
    }
    void brake() {
        print("Car is braking");
    }
} 
var names = <String>['Seth', 'Kathy', 'Lars'];
var uniqueNames = <String>{'Seth', 'Kathy', 'Lars'};
var pages = <String, String>{
  'index.html': 'Homepage',
  'robots.txt': 'Hints for web robots',
  'humans.txt': 'We are people, not machines'
};

class View {
  void build() {
    print('Building a view');
  }
}
var views = Map<String, View>();

void KwTest() {
  Object vh;
  try {
  }
  catch(e) {
    print(e);
  }
  finally {
    print("This is the end of the test");
  }

  var car = Car();
  car.steer();
  car.brake();

  final car2 = Car();
  car2.steer();
  car2.brake();

  Vehicle car3 = Car();
  car3.steer();
  car3.brake();
  
}


void main() {
  KwTest();
}