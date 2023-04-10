package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("Student Management", func() {
		var students map[string][]interface{}

		BeforeEach(func() {
			students = make(map[string][]interface{})
		})

		Describe("RemoveStudent", func() {
			It("should remove a student from the map", func() {
				students["John Doe"] = []interface{}{"123 Main St", "555-1234", 80}
				main.RemoveStudent(&students)("John Doe")
				Expect(len(students)).To(Equal(0))
				Expect(students["John Doe"]).To(BeNil())
			})
		})

		Describe("HighestScore", func() {
			It("should return the name and score of the student with the highest score", func() {
				students["John Doe"] = []interface{}{"123 Main St", "555-1234", 80}
				students["Jane Doe"] = []interface{}{"456 Main St", "555-5678", 90}
				students["Bob Smith"] = []interface{}{"789 Main St", "555-9012", 70}
				name, score := main.HighestScore(&students)
				Expect(name).To(Equal("Jane Doe"))
				Expect(score).To(Equal(90))
			})

			When("there are no students", func() {
				It("should return an empty string and 0", func() {
					name, score := main.HighestScore(&students)
					Expect(name).To(Equal(""))
					Expect(score).To(Equal(0))
				})
			})
		})
	})

})
