class BlloomFilter {
    // m
    private size: number;

    // k
    private hashCount: number;

    // 底层原理是字节（Byte），而不是位（Bit）
    // 存在巨大的空间浪费
    // 空间效率只有理论值的 1/8，造成了 8 倍的空间浪费
    // 数组的长度是 size / 8，因为每个元素可以存 8 个位（0 或 1）
    private bitArray: Uint8Array;

    
    /**
   * 构造一个高效的布隆过滤器
   * @param estimatedItemCount 预估要存储的元素数量 (n)
   * @param falsePositiveRate 期望的假阳性概率 (p)
   */
    constructor(estimatedItemCount: number = 1000, falsePositiveRate: number = 0.01) {
        const { size, hashCount } = this.getOptimalParams(estimatedItemCount, falsePositiveRate);
        this.size = size;
        this.hashCount = hashCount;
        const byteArraySize = Math.ceil(size / 8);
        this.bitArray = new Uint8Array(byteArraySize);
    }

    public add(item: string): void {
        const hashes = this.getHashes(item);
        for (const hash of hashes) {
            const byteIndex = Math.floor(hash / 8);
            const bitIndex = hash % 8;
            // 使用按位或 "|" 操作将该位置为 1
            this.bitArray[byteIndex] |= (1 << bitIndex);
        }
    }

    public mightContain(item: string): boolean {
        const hashes = this.getHashes(item);
        for (const hash of hashes) {
            const byteIndex = Math.floor(hash / 8);
            const bitIndex = hash % 8;
            // 使用按位与 "&" 操作检查该位是否为 1
            if ((this.bitArray[byteIndex] & (1 << bitIndex)) === 0) {
                return false;
            }
        }
        return true;
    }

    /**
     * 私有辅助方法：获取一个元素对应的 k 个哈希值（即位数组的索引）
     * 这里使用“双重哈希”技巧，仅用两个基础哈希函数来模拟 k 个哈希函数。
     * g_i(x) = (h1(x) + i * h2(x)) % m
     */
    private getHashes(item: string): number[] {
        const hashes: number[] = [];
        const h1 = this.djb2(item);
        const h2 = this.sdbm(item);

        for (let i = 0; i < this.hashCount; i++) {
            const hash = Math.abs((h1 + i * h2) % this.size);
            hashes.push(hash);
        }
        return hashes;
    }

    // --- 基础哈希函数 ---
    // 这些函数只需要速度快、分布好即可，不需要具备加密安全性。

    /**
     * 一个简单高效的字符串哈希函数 (djb2)
     */
    private djb2(str: string): number {
        let hash = 5381;
        for (let i = 0; i < str.length; i++) {
            hash = (hash * 33) ^ str.charCodeAt(i);
        }
        return hash;
    }

    /**
     * 一个简单的哈希函数 (sdbm)
     */
    private sdbm(str: string): number {
        let hash = 0;
        for (let i = 0; i < str.length; i++) {
            hash = str.charCodeAt(i) + (hash << 6) + (hash << 16) - hash;
        }
        return hash;
    }

    /**
   * 根据预估元素数量和期望误判率，计算最优的 m 和 k
   */
    private getOptimalParams(n: number, p: number): { size: number, hashCount: number } {
        const m = Math.ceil(-(n * Math.log(p)) / (Math.log(2) ** 2));
        const k = Math.round((m / n) * Math.log(2));
        return { size: m, hashCount: k };
    }
}