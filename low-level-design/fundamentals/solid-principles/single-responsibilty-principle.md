# Single Responsibilty Principle (SRP)
- A class should have single responsibility
- A class should have only 1 reason to change

**Goal** - This principle aims to separate behaviours so that if bugs arise as a result of your change, it won’t affect other unrelated behaviours.

![alt text](./resources/srp.png)

**Without SRP**
```c++
#include <iostream>
#include <fstream>
#include <string>

class Employee {
public:
    std::string name;
    double salary;

    Employee(const std::string& empName, double empSalary) 
        : name(empName), salary(empSalary) {}

    // This method is responsible for displaying employee details
    void displayEmployee() const {
        std::cout << "Name: " << name << ", Salary: $" << salary << "\n";
    }

    // This method is responsible for saving employee details to a file
    void saveToFile(const std::string& filename) const {
        std::ofstream file(filename);
        if (file.is_open()) {
            file << "Name: " << name << "\nSalary: " << salary << "\n";
            file.close();
        } else {
            std::cerr << "Error opening file\n";
        }
    }
};
```

**Issues:**
1. The Employee class handles two responsibilities:
   - Storing and displaying employee details.   
   - Saving details to a file.
2. If the file-saving logic changes (e.g., switching to a database), the Employee class needs modification, violating SRP.


**With SRP**
```c++
#include <iostream>
#include <fstream>
#include <string>

// Class responsible for storing and displaying employee details
class Employee {
public:
    std::string name;
    double salary;

    Employee(const std::string& empName, double empSalary) 
        : name(empName), salary(empSalary) {}

    void displayEmployee() const {
        std::cout << "Name: " << name << ", Salary: $" << salary << "\n";
    }
};

// Class responsible for saving employee data to a file
class EmployeeFileSaver {
public:
    void saveToFile(const Employee& emp, const std::string& filename) const {
        std::ofstream file(filename);
        if (file.is_open()) {
            file << "Name: " << emp.name << "\nSalary: " << emp.salary << "\n";
            file.close();
        } else {
            std::cerr << "Error opening file\n";
        }
    }
};
```

**Benefits:**
1. Separation of Concerns:
   - `Employee` handles only employee-related data.
   - `EmployeeFileSaver` handles file-saving logic.
2. Ease of Maintenance:
   - Changes to file-saving logic do not affect the `Employee` class.
3. Reusability: 
   - The `EmployeeFileSaver` class can be reused for different purposes without altering the `Employee` class.
