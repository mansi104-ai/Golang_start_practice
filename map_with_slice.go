package main

import "fmt"

// Define a Person struct with a name and a list of likes
type Person struct {
	Name  string
	Likes []string
}

// Function to group people by their likes
func groupPeopleByLikes() {
	// Initialize people with some sample data
	var people = []*Person{
		{Name: "Alice", Likes: []string{"Pizza", "Movies"}},
		{Name: "Bob", Likes: []string{"Pizza", "Soccer"}},
		{Name: "Charlie", Likes: []string{"Movies", "Soccer"}},
	}

	// Create a map to store what people like
	likes := make(map[string][]*Person)

	// Populate the map by iterating over people and their likes
	for _, p := range people {
		for _, l := range p.Likes {
			likes[l] = append(likes[l], p)
		}
	}

	// Print the grouped results
	for like, persons := range likes {
		fmt.Printf("People who like %s: ", like)
		for _, person := range persons {
			fmt.Printf("%s ", person.Name)
		}
		fmt.Println()
	}
}

func main() {
	// Call the function to group and display people by their likes
	groupPeopleByLikes()
}
