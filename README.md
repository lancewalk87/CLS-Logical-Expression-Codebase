# Logical Expression Codebase

#### ***This repository contains solutions to common problems, use cases and challenges within the field of computer science, and explores each of these solutions through multiple languages and methodologies.*** In addition to the language specific sources, the project is designed to be built from a CLI, provided you have each of the necessary binaries installed to your bash source, you should have no trouble running the application.

Running the CLI Application:
----------------------------

Project Structure:
------------------
* [**cli**](https://github.com/lancewalk87/CLS-Logical-Expression-Codebase/tree/master/cli): CLI Application Source.
  ```
  ├── bin                    
  ├── scr
  │   ├── build/
  │   └── system/
  └── src
  ```
* [**ent**](https://github.com/lancewalk87/CLS-Logical-Expression-Codebase/tree/master/ent): Language Build Sources.
  ```
  ├── input
  │   ├── *lang
  │   │   ├── main.*ext
  │   │   ├── problems/
  │   │   ├── services/
  │   │   └── usecases/
  └── output
  ```
* [**docs**](https://github.com/lancewalk87/CLS-Logical-Expression-Codebase/tree/master/ref): Project Documentation.
  ```

  ```

Expressed Languages:
-------------------

|     Language      | Ext.      | Attributes      | Source Shortcuts                                                                                               | Language References |
| ----------------- | --------- | --------------- | -------------------------------------------------------------------------------------------------------------- | ------------------- |
| C                 | .C        | GP, I, ST       | [ent/C](https://github.com/lancewalk87/CLS-Logical-Expression-Codebase/tree/master/ent/input/C)                      | [Jump To](https://www.cprogramming.com/)
| C-Plus-Plus       | .CPP      | GP, I, ST, OO   | [ent/C-Plus-Plus](https://github.com/lancewalk87/CLS-Logical-Expression-Codebase/tree/master/ent/input/C-Plus-Plus)  | [Jump To](http://www.cplusplus.com/)
| C-Sharp           | .CS       | GP, I, MP, OO   | [ent/C-Sharp](https://github.com/lancewalk87/CLS-Logical-Expression-Codebase/tree/master/ent/input/C-Sharp)          | [Jump To](https://docs.microsoft.com/en-us/dotnet/csharp/)
| Objective-C       | .M        | GP, OO          | [ent/Objective-C](https://github.com/lancewalk87/CLS-Logical-Expression-Codebase/tree/master/ent/input/Objective-C)  | [Jump To](https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ObjectiveC/Introduction/introObjectiveC.html)
| Swift             | .SWIFT    | GP, MP          | [ent/Swift](https://github.com/lancewalk87/CLS-Logical-Expression-Codebase/tree/master/ent/input/Swift)              | [Jump To](https://swift.org/)
| Java              | .JAVA     | GP, OO          | [ent/Java](https://github.com/lancewalk87/CLS-Logical-Expression-Codebase/tree/master/ent/input/Java)                | [Jump To](https://docs.oracle.com/javase/tutorial/)
| Kotlin            | .KT       | ST, GP, TI      | [ent/Kotlin](https://github.com/lancewalk87/CLS-Logical-Expression-Codebase/tree/master/ent/input/Kotlin)            | [Jump To](https://kotlinlang.org/)
| JavaScript        | .JS       | GP, MP          | [ent/JavaScript](https://github.com/lancewalk87/CLS-Logical-Expression-Codebase/tree/master/ent/input/JavaScript)    | [Jump To](https://www.javascript.com/)
| PHP               | .PHP      | SCR             | [ent/PHP](https://github.com/lancewalk87/CLS-Logical-Expression-Codebase/tree/master/ent/input/PHP)                  | [Jump To](http://www.php.net/)
| Python            | .PY       | GP              | [ent/Python](https://github.com/lancewalk87/CLS-Logical-Expression-Codebase/tree/master/ent/input/Python)            | [Jump To](https://www.python.org/)
| Ruby              | .RB       | GP, OO          | [ent/Ruby](https://github.com/lancewalk87/CLS-Logical-Expression-Codebase/tree/master/ent/input/Ruby)                | [Jump To](https://www.ruby-lang.org/en/)
| Go                | .GO       | GP, ST          | [ent/Go](https://github.com/lancewalk87/CLS-Logical-Expression-Codebase/tree/master/ent/input/Go)                    | [Jump To](https://golang.org/)
| Scala             | .SCALA    | GP, ST          | [ent/Scala](https://github.com/lancewalk87/CLS-Logical-Expression-Codebase/tree/master/ent/input/Scala)              | [Jump To](https://www.scala-lang.org/)
| Lua               | .LUA      | MP, LW          | [ent/Lua](https://github.com/lancewalk87/CLS-Logical-Expression-Codebase/tree/master/ent/input/Lua)                  | [Jump To](https://www.lua.org/)
| Rust              | .RS       | MP              | [ent/Rust](https://github.com/lancewalk87/CLS-Logical-Expression-Codebase/tree/master/ent/input/Rust)                | [Jump To](https://www.rust-lang.org/)
| Perl              | .PERL     | GP              | [ent/Perl](https://github.com/lancewalk87/CLS-Logical-Expression-Codebase/tree/master/ent/input/Perl)                | [Jump To](https://www.perl.org/)

| Shorthand | Definition         |
| --------- | ------------------ |
| I         | Imperative         |
| GP        | General Purpose    |
| ST        | Statically Types   |
| WT        | Weakly Typed       |  
| OO        | Object Oriented    |   
| MP        | Multi-Paradigm     |  
| TI        | Type-Inference     |
| SCR       | Scripting Lang.    |  
| LW        | Light Weight       |  


Ubiquitously Applied Algorithms:  
--------------------------------
| Name        | Average Complexity    | Worst Complexity | Basic Definition                                      |
| ----------- | --------------------- | ---------------- | ----------------------------------------------------- |
| Selection   | n^2                   | n^2              | O time complex                                        |
| Insertion   | n^2                   | n^2              | product build one increment per operation             |
| Bubble      | n^2                   | n^2              | swap pairs in list upon full incrementation of input  |
| Quick       | n*log(n)              | n^2              |  O efficient, ordering list                           |
| Heap        | n*log(n)              | n*log(n)         | sorts heap binary tree by comparing size of max input |   
| Merge       | n*log(n)              | n*log(n)         | divides complex input to smaller subset problems      |
| Shell       | n*log(n)^2 or n^(3/2) | n/a              | complex application of general sort algorithm         |
| Topological | n/a                   | n/a              | linear order of input vertices                        |


Solved Problems
--------
* [Structure and Interpretation of Computer Programs: ](https://github.com/lancewalk87/CLS-Logical-Expression-Codebase/tree/master/docs/structure_interpretation.md) Fundamental principles of Computer Science.  
* [An Introduction to Algorithms: ](https://github.com/lancewalk87/CLS-Logical-Expression-Codebase/tree/master/docs/introduction_algorithms.md) General introduction to algorithms and their applications.
* [Concrete mathematics: ](https://github.com/lancewalk87/CLS-Logical-Expression-Codebase/tree/master/docs/concrete_mathematics.md) Concepts of Computer Science from a mathematical perspective.
* [Project Euler: ](https://github.com/lancewalk87/CLS-Logical-Expression-Codebase/tree/master/docs/project_euler.md) A website dedicated to the fascinating world of mathematics and programming.

Notable References, Documentation, and Manuals:
-----------------------------------------------
