{

    if ($0 ~ /([aeiou].*){3,}/ && $0 !~ /(ab|cd|pq|xy)/) print NR " " $0
}