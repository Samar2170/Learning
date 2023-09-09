class Point {
  double? x;
  double? y;
  double z = 0;
}

void a() {
  var p = Point();
  p.x = 4;
  p.y = 3;
  p.z = 2;
  print(p.x);
  print(p.y);
  print(p.z);
}

class ProfileMark {
  final String name;
  final DateTime start = DateTime.now();

  ProfileMark(this.name);
  ProfileMark.unnamed() : name = '';
}

void b() {}
