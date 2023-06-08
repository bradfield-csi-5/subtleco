# https://leetcode.com/problems/course-schedule/
from collections import defaultdict


class Solution:
    def __init__(self):
        self.courses = set()
        self.cleared_courses = set()
        self.unseen_courses = set()
        self.graph = defaultdict(list)

    def clear_course(self, course):
        self.cleared_courses.add(course)
        self.unseen_courses = self.courses.difference(self.cleared_courses)

    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        # check for no preqs
        if not prerequisites:
            return True

        # build a graph
        for preq in prerequisites:
            course = preq[0]
            prerequisite = preq[1]

            self.courses.add(course)
            self.courses.add(prerequisite)
            self.graph[course].append(prerequisite)

        # find base case of classes with no prerequisite
        for course in self.courses:
            if not self.graph[course]:
                self.clear_course(course)

        # kill it with fire
        if not self.cleared_courses:
            return False

        for course in self.unseen_courses:
            self.check_course(numCourses, course)

        return len(self.cleared_courses) >= numCourses

    def check_course(self, numCourses, course):
        if len(self.cleared_courses) >= numCourses:
            return true

        # Are all the preqs cleared?
        num_of_preqs = len(self.graph[course])
        cleared_preqs = 0
        for preq in self.graph[course]:
            if preq in self.cleared_courses:
                cleared_preqs += 1
        if num_of_preqs == cleared_preqs:
            self.clear_course(course)
