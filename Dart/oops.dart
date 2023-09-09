import 'classes.dart';

class ClassName {
  int a;
  int b;
  ClassName(this.a, this.b);

  void display() {
    print(a);
    print(b);
  }
}

// Encapsulation
// using underscore _ for files and classes both

// Getter and Setter Methods
// used to update private properties

class Employee {
  String _name;

  String getName() => _name;

  void setName(String name) {
    this._name = name;
  }

  Employee(this._name);
}

// Inheritance

class ParentClass {}

class ChildClass extends ParentClass {}

// Inheritance of constructor

class Laptop {
  String? name;
  String? color;

  String? get getName => name;
  String get fullName => "$name $color";
  Laptop(String name, String color) {
    print("Laptop constructor");
    print("Name: $name");
    print("Color: $color");
    this.name = name;
    this.color = color;
  }
}

class Macbook extends Laptop {
  final String arch;
  Macbook({required String name, required String color, required this.arch})
      : super(name, color);

  String get fullName => "$name $arch";
}

void testInheritance() {
  var macbook = Macbook(name: "Macbook Pro", color: "Space Grey", arch: "M1");
  var laptop = macbook as Laptop;
  print(laptop.name);
  print(laptop.fullName);
}

//  Static vars, accessible by all instances of a class.

class Vehicle {
  static String brand = 'Maruti';
}

// mixins

mixin Electric {
  void electricVariant() {
    print('This is an electric variant');
  }
}

mixin Petrol {
  void petrolVariant() {
    print('This is a petrol variant');
  }
}

class Car with Electric, Petrol {}

void testMixins() {
  var car = Car();
  car.electricVariant();
  car.petrolVariant();
}

// generics
class Data<T> {
  T data;
  Data(this.data);
}

void testGenerics() {
  Data<int> intData = Data<int>(10);
  Data<double> doubleData = Data<double>(10.5);

  // print the data
  print("IntData: ${intData.data}");
  print("DoubleData: ${doubleData.data}");
}
