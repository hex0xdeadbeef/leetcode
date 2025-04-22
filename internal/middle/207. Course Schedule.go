package middle


func canFinish(numCourses int, prerequisites [][]int) bool {
    graph := make([][]int, numCourses)
    for _, p := range prerequisites {
        graph[p[1]] = append(graph[p[1]], p[0])
    }

    status := make([]int, numCourses)

    var dfs func(int) bool
    dfs = func(course int) bool {
        if status[course] == 1 {
            return false
        }
        if status[course] == 2 {
            return true
        }

        status[course] = 1
        for _, next := range graph[course] {
            if !dfs(next) {
                return false
            }
        }
        status[course] = 2
        return true
    }

    for c := range numCourses {
        if status[c] == 0 {
            if !dfs(c) {
                return false
            }
        }
    }
    return true
}
