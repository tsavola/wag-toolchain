extern "C" {
	void __start();
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
	return n * 0x10001;
}

void __start()
{
	set_result(main());
}
