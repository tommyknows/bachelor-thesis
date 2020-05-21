public static void listLambdas(List<Integer> list) {
    // Sammelt alle Elemente der Liste um 1 inkrementiert
    // in einer neuen Liste
    List <Integer > incremented = new ArrayList<Integer>();
    for(int n : list) {
        incremented.add(n+1);
    }

    // Variante mit Lambdas
    List<Integer> incremented2 = list.stream()
        .map(x -> x + 1)
        .collect(Collectors.toList());

    // Skaliert die Element in der Liste um den Faktor s
    // und addiert dazu jeweils t;
    // sammelt das Ergebnis in einer neuen Liste
    final int s = 2;
    final int t = 1;
    List<Integer> scaled = list.stream()
        .map(x -> s * x + t)
        .collect(Collectors.toList());

    // Filtert alle geraden Elemente aus der Liste
    // und sammelt sie in einer neuen Liste
    List<Integer> even = list.stream()
        .filter(x -> x % 2 == 0)
        .collect(Collectors.toList());

    // Summiert die Elemente in der Liste
    int summe_reduce = 0;
    for(int n : list) {
        summe_reduce += n;
    }

    // Variante mit Lambdas
    int summe_reduce2 = list.stream()
        .reduce(0, (a, b) -> a + b);
}
