extern "C" {
	int main();
	void putns(const char *str, int len);
}

static void barrier()
{
	putns("barrier", 7);
}

int n;

int main()
{
	n |= 1 << 4;
	barrier();
	n |= 1 << 8;
	barrier();
	n |= 10 << 12;
	return n * 0x10001;
}
