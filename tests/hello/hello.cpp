extern "C" {
	void set_result(int retval);
}

int n;

int main()
{
	n |= 1 << 4;
	set_result(0); // barrier
	n |= 1 << 8;
	set_result(0); // barrier
	n |= 10 << 12;
	set_result(n * 0x10001);
	return 0;
}
