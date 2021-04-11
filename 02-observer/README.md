# Observer Pattern

## 특징

1. 신문 구독 매커니즘과 같다. 출판사를 subject, 구독자를 observer라고 부른다.
2. 한 객체 상태가 바뀌면 그 객체에 의존하는 다른 객체들에게 연락이 간다. 자동으로 데이터가 업데이트되는 방식으로 일대다 의존성을 정의한다.
3. 두 객체가 느슨하게 결합되어 있다는 것은 상호작용을 하지만 서로 모른다는 의미이다. 옵저버 패턴에서 subject와 observer는 느슨하게 결합되어 있는 객체 디자인이다.
4. subject가 observer에 대해 아는 것은 observer가 observer interface를 구현한다는 것이다.
5. observer는 추가, 제거가 자유롭다.
6. 새로운 observer를 등록해도 subject는 변하지 않는다.
7. subject와 observer는 서로 독립적이다.

![class diagram](https://2.bp.blogspot.com/-xgiuTvAD4EI/Wy4qkZJDHmI/AAAAAAAACj4/xVrxGOVR2V452XUKain8m-UOTlxxGuJBgCLcBGAs/s640/observer-generic-class-diagram.PNG)

![implements example](https://4.bp.blogspot.com/-9RAV8APWNxw/Wy4s9nx5TtI/AAAAAAAACkc/yQaRxbUEu1IdqOYF8Fk6VRsUrm8DEY1ygCLcBGAs/s640/observer-weatherstation-class-diagram.PNG)
