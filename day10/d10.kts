
data class State(val items: List<Int>, val pos: Int, val skip: Int)

fun doMove(currState: State, pinchSize: Int): State {
    val pinched = (currState.items + currState.items)
        .slice(currState.pos until (currState.pos + pinchSize))
        .reversed()

    val nextList = currState.items.toMutableList()
    pinched.withIndex().forEach { (idx, item) ->
        nextList[(idx + currState.pos) % nextList.size] = item
    }

    return State(
            nextList,
            (currState.pos + pinchSize + currState.skip) % nextList.size,
            currState.skip + 1)
}

fun evalMoves(initialState: State, lengths: List<Int>): State {
    var state = initialState
    lengths.forEach { state = doMove(state, it) }
    return state
}

fun p1() {
    val testResult = evalMoves(
            State((0..4).toList(), 0, 0),
            listOf(3, 4, 1, 5)).items
    println(testResult[0] * testResult[1])

    val inputMoves = listOf(189,1,111,246,254,2,0,120,215,93,255,50,84,15,94,62)
    val p1Result = evalMoves(
            State((0..255).toList(), 0, 0),
            inputMoves).items
    println(p1Result[0] * p1Result[1])
}

p1()


fun strToAscii(str: String): List<Int> = str.map { it.toInt() }

fun makeDense(sparse: List<Int>) =
    sparse.chunked(16) { it.reduce(Int::xor) }

fun solveP2(input: String) {
    val extraLengths = listOf(17, 31, 73, 47, 23)
    val nRounds = 64
    val lengths = strToAscii(input) + extraLengths
    var state = State((0..255).toList(), 0, 0)
    (0 until nRounds).forEach { state = evalMoves(state, lengths) }
    val dense = makeDense(state.items)
    println(dense.map { num ->
        val hexStr = num.toString(16)
        if (hexStr.length == 1) {
            "0$hexStr"
        } else {
            hexStr
        }
    }.joinToString(""))
}

fun p2() {
    solveP2("")
    solveP2("AoC 2017")
    solveP2("1,2,3")
    solveP2("1,2,4")
    val input = "189,1,111,246,254,2,0,120,215,93,255,50,84,15,94,62"
    solveP2(input)
}

p2()
