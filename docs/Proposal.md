

# Critterarium Proposal:

**Team Members:**

* Jack Werremeyer
* Nolan


## 1. Project Summary

Critterarium is a distributed simulation where every virtual critter is controlled by a small custom programming language called CritterLang. Each critter runs its own CritterLang program, which determines how it moves, eats, reacts to neighbors, and reproduces inside a shared simulated world.

Hopefully, by the end of the project, the system will resemble an ecosystem, where mutations and interactions can result in fun and interesting behaviors


## 2. Typical Use Cases

1. Programming Critters

The user writes a CritterLang program that defines how a fish behaves and the program is parsed, validated, and interpreted by the simulation engine.

2. Running the ecosystem

The backend server simulates the aquarium world, running thousands of critter programs concurrently. Critters interact with each other and the environment in real time.


3. GUI Monitoring and Control

Users connect to the simulation using a GUI client to visualize critters, pause or resume the simulation, inspect individual programs, and adjust parameters like mutation rate or food density. (maybe count predator too)



## 3. Intended Design and Components

The project will be organized into multiple components:

### Core Modules

1. critterlang/token


   * Converts source code into tokens

2. critterlang/parser

   * Parses tokens into an Abstract Syntax Tree

3. critterlang/ast

   * Abstract Syntax Tree defs
   * Used parser, interpreter, and fault injector

4. critterlang/interpreter

   * Executes commands for individual critters
   * Maintain state for critter

5. critterlang/prettyprinter

   * Given an Abstract Syntax Tree, create a formatted CritterLang program

6. critterlang/faultinjector

   * Randomly mutates tree nodes
   * Used to simulate genetic mutation during reproduction



### Simulation/System Modules

7. simulation/world

   * Aquarium grid or continuous space that handles physics, food spawning, collisions, and interactions

8. simulation/server

   * Multi-threaded Go server that runs the simulation loop and critter interpreters

10. client/gui

    * GUI for visualizing the aquarium and controlling the simulation


## 4. Testing Strategy

Testing will focus on parsing/interpretating the CritterLang correctly, ensuring mutations works properly, and the simulation run without a hitch



## 5. Minimal Viable Product 

The MVP will include:

* Working CritterLang parser that builds an Abstract Syntax Tree
* A basic interpreter that can execute simple critter behaviors
* A fault injector
* A command-line/minimal visualization of the simulation
* Necessary Unit tests for parser/interpreter/fault injector



## 6. Stretch Goals

* A GUI visualization (maybe aquarium themed?)
* Distributed simulation across multiple servers
* More advanced CritterLang features 


## 7. Expected Functionality by Checkpoint

By the checkpoint, the project is expected to include:

* Completed parser
* Fully defined Abstract syntax tree 
* Working pretty printer
* Unit tests
* Basic scaffolding for the simulation server
