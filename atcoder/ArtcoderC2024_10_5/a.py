def min_max_lunch_group(N, K):
    """
    キーエンス本社の従業員を昼休みの時間帯で2つのグループに分ける際に、
    同時に昼休みをとる人数の最大値を最小にする関数

    Args:
        N: 部署の数
        K: 各部署の人数を表すリスト

    Returns:
        同時に昼休みをとる最大人数としてあり得る最小の値
    """
    INF = float('inf')
    total = sum(K)
    ans = INF

    # ビット全探索
    for i in range(1 << N):
        group_a = 0
        for j in range(N):
            if i & (1 << j):
                group_a += K[j]
        group_b = total - group_a
        ans = min(ans, max(group_a, group_b))

    return ans


# 入力
if __name__ == "__main__":
    N = int(input().strip())
    K = list(map(int, input().strip().split()))
    if len(K) != N:
        raise ValueError("The number of departments does not match the provided list length.")
    
    # 計算 & 出力
    result = min_max_lunch_group(N, K)
    print(result)