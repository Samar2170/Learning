void vars() {
    String Name = "John";
  int Age = 35;
  String sentence = Printer(Name, Age);
  print(sentence);

  int num1 = 10;
  int num2 = 20;
  int num3 = addNumbers(num1, num2);
  print(num3);

  String ans = Conditionals(false);
  print(ans);

  looping();

}
int addNumbers(int num1, int num2) {
  return num1 + num2;
}


String Printer(String name, int age) {
  return "My name is $name and I am $age years old";
}

String Conditionals(bool isSmart) {
  if(!isSmart){
    return "You are not smart";
  } else {
    return "You are smart";
  }
}

void looping() {
  int i = 0;

  while(i<5) {
    print(i);
    i++;
  }

  for (int i =0; i<10; i++) {
    print(i);
  }
}