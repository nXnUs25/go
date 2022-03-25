# Questions 1
---

## Questions 

- ##### 1. What is the definition of software engineering? 
- ##### 2. What are some of the questions that every SWE should be able to answer? 
- ##### 3. Compare the role of an SWE and an SRE. What are the key differences between the two roles? 
- ##### 4. Name some of the deficiencies of the waterfall model. Explain how the iterative enhancement model attempts to address those deficiencies. 
- ##### 5. What are the most common sources of waste according to the lean development model? 
- ##### 6. Provide an example where focusing all the optimization efforts on a single step of the development process can have a negative effect on the efficiency of the end-to-end process. 
- ##### 7. What are the key responsibilities of the PO and the SM in the Scrum framework? 
- ##### 8. What is the role of retrospectives in Scrum? What topics should the team be discussing and what should be the expected outcome of each retrospective session? 
- ##### 9. Why are automation and measuring important when following the DevOps model? 
- ##### 10. You are working for ACME Gaming System, a company with the vision to disrupt the already mature and highly competitive gaming console market. To this end, the company has partnered with Advanced GPU Devices, a well-known graphics chip manufacturer, to prototype a new GPU design that should allow the upcoming gaming console to blow competitor consoles out of the water. Your task, as the project's lead engineer, is to design and build the software that will power the new console. Which software development model would you choose? Explain the reasoning behind your decision.

---
# _NOTE:**_ 
### I wasnt answering any kind of questions from first section, but is worth to read the answers provided by author and keep it as notes.
---
# Answers to Chapter 1 ... by author of the book 

### Chapter 1 

#### 1. 
Software engineering is defined as the application of a systematic, disciplined, quantifiable approach to the development, operation, and maintenance of software. 

#### 2. 
Some of the key questions that a software engineer must be able to answer are as follows: 
- What are the business use cases that the software needs to support? 
- What components comprise the system and how do they interact with each other? 
- What technologies will be used to implement the various system components? 
- How will the software be tested to ensure that its behavior matches the customer's expectations? 
- How does load affect the system's performance and what is the plan for scaling the system? 

#### 3. 
An SRE spends approximately half of their time on operations-related tasks such as dealing with support tickets, being on call, automating processes to eliminate human errors, and so on. 

#### 4. 
The _waterfall model_ does not provide a detailed view of the processes that comprise each model step. In addition, it does not seem to support _cross-cutting processes_ such as project management or quality control that run in parallel with the waterfall steps. A significant caveat of the waterfall model is that it operates under the assumption that all customer requirements are known in advance. The iterative enhancement model attempts to rectify these issues by executing small incremental waterfall iterations that allow the development team to adapt to changes in requirements.
 
#### 5.
According to the lean development model, the most common sources of waste are as follows: The introduction of non-essential changes when development is underway Overly complicated decision-making processes for signing off new features Unneeded communication between the various project stakeholders and the development teams 
 
#### 6. 
The team decides to focus on speedy delivery at the expense of code quality. As a result, the code base becomes more complex and defects start accumulating. Now, the team must dedicate a part of their development time to fixing bugs instead of working on the requested features. Consequently, the implementation stage becomes a bottleneck that reduces the efficiency of the entire development process. 

#### 7. 
The key responsibility of the Product Owner is to manage the backlog for the project. On the other hand, the Scrum Master is responsible for organizing and running the various Scrum events. 

#### 8. 
The retrospective serves as a feedback loop for incrementally improving the team's performance across sprints. The team members should be discussing both the things that went well during the last sprint as well as the things that didn't. The outcome of the retrospective should be a list of corrective actions to address the problems that were encountered during the sprint. 

#### 9. 
Automation is important as it reduces the potential for human error. In addition, it reduces the time that's needed to test and deploy changes to production. Measuring is equally important as it allows DevOps engineers to monitor production services and receive alerts when their behavior diverges from the expected norm. 

#### 10. 
The company is expected to operate in a high-risk environment. For one, the new gaming console depends on a piece of technology that is not available yet and is being developed by a third party. What's more, the market is already saturated: other, much larger competitors could also be working on their own next-gen console systems. The expected competitive advantage for ACME Gaming Systems may be rendered obsolete by the time their new system is released. This is yet another source of risk. Given the high risk that's involved, the spiral model with its risk assessment and prototyping processes would be the most sensible choice for developing the software that will power the new console.