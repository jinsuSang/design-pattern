# Factory Pattern

## 특징

1. 객체를 생성하기 위한 인터페이스를 정의한다. 어떤 클래스의 인스턴스를 만들지는 서브클래스에서 결정한다.
2. 클래스의 인스턴스를 만드는 일을 서브클래스에게 맡긴다.
3. 객체 생성 코드를 전부 한 객체 또는 메소드에 집어 넣어 코드 중복을 제거한다.
4. "의존성 뒤집기 원칙" : 추상화 된 것에 의존하도록 한다. 구상 클래스에 의존하도록 만들지 않도록 한다.

### "의존성 뒤집기 원칙"

- 어떤 변수에도 구상 클래스에 대한 레퍼런스를 저장하지 않는다.
- 구상 클래스에서 유도된 클래스를 만들지 않는다
- 베이스 클래스에 이미 구현되어 있던 메소드를 오버라이드하지 않는다.

### 추상 팩토리 패턴

1. 인터페이스를 이용하여 서로 연관된 또는 의존하는 객체를 구상 클래스를 지정하지 않고 생성할 수 있다.

## 팩토리 메소드 패턴 VS 추상 팩토리 패턴

1. 팩토리 메소드 패턴은 상속으로 객체를 형성, 추상 팩토리 패턴은 객체 구성(Composition)으로 형성
2. 팩토리 메소드 패턴으로 객체를 생성할 때는 클래스를 확장하고 맥토리 메소드를 오버라이드 해야한다.
3. 추상 팩토리 패턴은 팩토리를 만들 때 먼저 인스턴스를 만들고 추상 형식을 써서 만든 코드에 전달한다.
4. 추상 팩토리 패턴은 새로운 제품이 추가될 시 인터페이스를 변경해야 한다.
5. 둘 모두 객체 생성을 캡슐화해서 애플리케이션 결합을 느슨하게 만들고 특정 구현에 덜 의존적으로 한다.

---

## Simple Factory Pattern

![simple factory](https://www.oreilly.com/library/view/head-first-design/0596007124/figs/web/119fig01.png.jpg)

## Factory Pattern

![factory pattern](https://www.oreilly.com/library/view/head-first-design/0596007124/figs/web/134fig01.png.jpg)

## Abstract Factory Pattern

![abstract factory pattern diagram](https://www.oreilly.com/library/view/head-first-design/0596007124/figs/web/158fig01.png.jpg)
![abstract factory pattern](https://www.oreilly.com/library/view/head-first-design/0596007124/figs/web/159fig01.png.jpg)

## Dependency Inversion Principle Example

![depencency](https://www.oreilly.com/library/view/head-first-design/0596007124/figs/web/140fig01.png.jpg)

![inversion](https://www.oreilly.com/library/view/head-first-design/0596007124/figs/web/142fig01.png.jpg)
