import java.io.File

fun addPipe(pipes: MutableMap<Int, Set<Int>>, pipe: Pair<Int, List<Int>>) {
    val (lhs, rhs) = pipe
    val currLhsVal = pipes.getOrElse(lhs, { setOf<Int>() })
    pipes[lhs] = currLhsVal union rhs.toSet()
    rhs.forEach { proc ->
        val currRhsVal = pipes.getOrElse(proc, { setOf<Int>() })
        pipes[proc] = currRhsVal union setOf<Int>(lhs)
    }
}

fun parseInput(pipes: String): Map<Int, Set<Int>> {
    val parsedPipes = pipes.lines().map { line ->
        val (lhs, rhs) = line.split("<->").map(String::trim)
        Pair(
                lhs.toInt(),
                rhs.split(",").map(String::trim).map(String::toInt))
    }
    val output = mutableMapOf<Int, Set<Int>>()
    parsedPipes.forEach { addPipe(output, it) }
    return output
}

fun traceFromID(pipes: Map<Int, Set<Int>>, initial: Int): Set<Int> {
    var visited = setOf(initial)
    val queue = mutableListOf<Int>()
    queue.add(initial)
    while (queue.isNotEmpty()) {
        val next = queue.removeAt(0)
        val nextLinks = pipes.getOrElse(next, { setOf<Int>() })
        queue.addAll(nextLinks - visited)
        visited = visited union nextLinks
    }
    return visited
}

fun countFromID(pipes: Map<Int, Set<Int>>, initial: Int): Int {
    return traceFromID(pipes, initial).size
}

fun countNumGroups(pipes: Map<Int, Set<Int>>): Int {
    var remaining = pipes.keys
    var count = 0
    while (remaining.size > 0) {
        count++
        val next = remaining.first()
        val reachable = traceFromID(pipes, next)
        remaining = remaining - reachable
    }
    return count
}


val testInput = """
0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5""".trim()

val input = File("./input.txt").readText().trim()

fun solveP1() {
    val pipes = parseInput(testInput)
    println(countFromID(pipes, 0))
    println(countFromID(parseInput(input), 0))
}

solveP1()

fun solveP2() {
    println(countNumGroups(parseInput(testInput)))
    println(countNumGroups(parseInput(input)))
}

solveP2()
