# Students Records Manager

## Objectives
Develop a simplified Student Records Management System in Go. The system should be able to read student records from a file, process them, and write the processed information back to a file.

## Instructions
### Datas structure
- You have to create a `Student` structure able to contains thoses informations :  
  - Name
  - Age
  - Grade
  - ID
- Create a `School` structure that contain a **list of Students**  
  
### Files manipulation
- Write a function to **load a school** from a CSV file, parse the file and return an instance of `School`. If the file doesn't exist, return an error.  
This function **must have a filename parameter**.  
- Write a function to **save** the current state of the school in a CSV file and return an error if the file was didn't save correctly.

### Features
- Write a function to calcultate and return the average grade of students in the school.
- Write a function to find and return the best student of the school.
- Write a function to return a list students order by descending grade.

## Evaluation
- Features
- The readability and cleanliness of the code
- Error management
- Ability to search for information and read documentation

## Contraints
You are not allowed to use ChatGPT or others IA tool to make this exercice.