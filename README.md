# SINGLETON PATTERN
Usually, a singleton instance is created when the struct is first initialized. To make this happen, we define the getInstance method on the struct.
This method will be responsible for creating and returning the singleton instance. Once created, the same singleton instance will be returned every time the getInstance is called.

- There is a nil-check at the start for making sure singleInstance is empty first time around. This is to prevent expensive lock operations every time the getinstance method is called.
If this check fails, then it means that the singleInstance field is already populated.

- The singleInstance struct is created within the lock.

- The sync.Once will only perform the operation once.

# STRATEGY PATTERN
- BowSolider: a soldier with a BOW.
- ShieldSoldier: a soldier with a SHIELD.
- ShieldBowSoldier: a soldier that picked up a SHIELD and a BOW.

The problem is how to decouple our soldier class from these weapons(algorithms) so that we can change the weapon(algorithm) at run time. Also, the soldier class should not change when a new weapon(algorithm) is being added.

This is where Strategy pattern comes into the picture. It suggests creating a family of the algorithm with each algorithm having its own class. Each of these classes follows the same interface, and this makes the algorithm interchangeable within the family.
Let’s say the common interface name is ISoldier.

Now our main soldier class will embed the ISoldier interface. Instead of implementing all types of weapon's algorithms in itself, our soldier class will delegate the execution to the ISoldier interface.
Since ISoldier is an interface, we can change the algorithm in run time to either BowSoldier, ShieldSoldier, ShieldBowSoldier without changing the soldier class.

# DECORATOR PATTERN
Using decorators you can wrap objects countless number of times since both target objects and decorators follow the same interface. The resulting object will get a stacking behavior of all wrappers.

We have basicSoldier with Attack=10, HP=100, at some point our basicSoldier found a bow, now he is bowSoldier with Attack=10+15, HP=100-20, after that he desides to took a shield also, so we are going to wrap wrapped basicSoldier, and get Attack=10+15-5, HP=100-20+50.

# ADAPTER PATTERN
The Adapter acts as a wrapper between two objects. It catches calls for one object and transforms them to format and interface recognizable by the second object.

- We have a adapter code that expects some features of an object (Greek speech), but we have another object called translator (Roman translator) which offers the same functionality but through a different interface (Roman speech).
- This is where the Adapter pattern comes into the picture. We create a struct type known as roman translator that will: Adhere to the same interface which the adapter expects (Greek speech).
- Translate the request from the adapter to the translator in the form that the adapter expects.
- The translator accepts a greek speech and then translates its signals into a roman speech and passes them to the speakingRoman in RomanSoldier.

# FACTORY PATTERN
It’s impossible to implement the classic Factory Method pattern in Go due to lack of OOP features such as classes and inheritance.
However, we can still implement the basic version of the pattern, the Simple Factory.

In this example, we’re going to create various types of soldiers using a factory struct.

- First, we create the ISoldier interface, which defines all methods a soldier should have.
- There is a basicSoldier struct type that implements the ISoldier interface. Two concrete BowSoldier and ShieldSoldier embed solder struct and indirectly implement all ISoldier methods.
- The soldierFactory struct serves as a factory, which creates soldiers of the desired type based on an incoming argument. The main.go acts as a client. Instead of directly interacting with Archer or Shieldbearer, it relies on soldierFactory to create instances of various soldiers, only using string parameters to control the production.

# OBSERVER PATTERN
The Observer pattern provides a way to subscribe and unsubscribe to and from these events for any object that implements a subscriber interface.

In our game e-commerce shop, items go out of stock from time to time. There can be customers who are interested in a particular item that went out of stock. There are three solutions to this problem:

The customer keeps checking the availability of the item at some frequency.
E-commerce shop bombards customers with all new items available, which are in stock.
The customer subscribes only to the particular item he is interested in and gets notified if the item is available. Also, multiple customers can subscribe to the same product.
Option 3 is most viable, and this is what the Observer pattern is all about. The major components of the observer pattern are:

Subject, the instance which publishes an event when anything happens.
Observer, which subscribes to the subject events and gets notified when they happen.
