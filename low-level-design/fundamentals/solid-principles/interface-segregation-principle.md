Interface Segregation Principle

- Interface should be such that client should implement unnecessary functions they do not need.
- Clients should not be forced to depend on methods that they do not use.

**Goal:** This principle aims at splitting a set of actions into smaller sets so that a Class executes ONLY the set of actions it requires.

This principle focuses on creating smaller, more specific interfaces rather than having one large, all-encompassing interface. By doing this, we ensure that classes only implement methods relevant to their functionality, improving modularity and flexibility.

![alt text](./resources/isp.png)

**Without ISP:**
```c++
#include <iostream>
#include <string>

// Single interface for all printers
class Printer {
public:
    virtual void printDocument(const std::string& doc) = 0;
    virtual void scanDocument(const std::string& doc) = 0;
    virtual void faxDocument(const std::string& doc) = 0;
};

class BasicPrinter : public Printer {
public:
    void printDocument(const std::string& doc) override {
        std::cout << "Printing: " << doc << "\n";
    }

    void scanDocument(const std::string& doc) override {
        throw std::logic_error("Scan not supported!");
    }

    void faxDocument(const std::string& doc) override {
        throw std::logic_error("Fax not supported!");
    }
};
```

**Issues:**
1. The `BasicPrinter` class is forced to implement methods for scanning and faxing, which it does not support.
2. This violates ISP because `BasicPrinter` is dependent on methods it doesn't use.
3. The design introduces runtime errors when unsupported methods are called.


**With ISP:**
```c++
#include <iostream>
#include <string>

// Specific interfaces
class Printable {
public:
    virtual void printDocument(const std::string& doc) = 0;
};

class Scannable {
public:
    virtual void scanDocument(const std::string& doc) = 0;
};

class Faxable {
public:
    virtual void faxDocument(const std::string& doc) = 0;
};

// Basic printer implements only the Printable interface
class BasicPrinter : public Printable {
public:
    void printDocument(const std::string& doc) override {
        std::cout << "Printing: " << doc << "\n";
    }
};

// Advanced printer supports printing and scanning
class AdvancedPrinter : public Printable, public Scannable {
public:
    void printDocument(const std::string& doc) override {
        std::cout << "Printing: " << doc << "\n";
    }

    void scanDocument(const std::string& doc) override {
        std::cout << "Scanning: " << doc << "\n";
    }
};

// All-in-one printer supports printing, scanning, and faxing
class AllInOnePrinter : public Printable, public Scannable, public Faxable {
public:
    void printDocument(const std::string& doc) override {
        std::cout << "Printing: " << doc << "\n";
    }

    void scanDocument(const std::string& doc) override {
        std::cout << "Scanning: " << doc << "\n";
    }

    void faxDocument(const std::string& doc) override {
        std::cout << "Faxing: " << doc << "\n";
    }
};
```

**Benefits of ISP:**
1. Focused Interfaces:
   1. Classes implement only the methods they need. For example, `BasicPrinter` doesn't need to implement `scanDocument` or `faxDocument`.
2. Extensibility:
   1. New functionality can be added by creating new interfaces without affecting existing classes.
3. Avoidance of Unused Methods:
   1. Classes are not burdened with irrelevant methods, improving readability and maintainability.