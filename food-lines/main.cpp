//
// Created by Aybars Nazlica on 2025/03/02.
//

#include <iostream>
#include <vector>


int find_shortest_line_index(const std::vector<int>& lines)
{
    int shortest_line_index = 0;
    for (int i = 0; i < lines.size(); i++)
    {
        if (lines[i] < lines[shortest_line_index])
        {
            shortest_line_index = i;
        }
    }
    return shortest_line_index;
}

void solve(std::vector<int>& lines, const int m)
{
    for (int i = 0; i < m; i++)
    {
        const int shortest_line_index = find_shortest_line_index(lines);
        std::cout << lines[shortest_line_index] << '\n';
        lines[shortest_line_index]++;
    }
}

int main()
{
    int n, m;
    std::cin >> n >> m;
    std::vector<int> lines(n);

    for (int& line : lines)
    {
        std::cin >> line;
    }

    solve(lines, m);

    return 0;
}
