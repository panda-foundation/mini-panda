#include <stdint.h>

struct test_Pwm;

struct test_Cpu;

struct test_Gpu;

void main();

void test_expression();

void test_test_unary();

void test_test_increment();

void test_test_decrement();

void test_test_binary();

void test_test_others();

void test_test_initializer();

void test_test_subscripting();

void test_test_scope();

void test_test_member_access();

void test_test_conversion();

void global_assert(uint8_t expression, uint8_t* message);

void test_statement();

void console_write_bool(uint8_t value);

void console_write_i8(int8_t value);

void console_write_i16(int16_t value);

void console_write_i32(int32_t value);

void console_write_i64(int64_t value);

void console_write_u8(uint8_t value);

void console_write_u16(uint16_t value);

void console_write_u32(uint32_t value);

void console_write_u64(uint64_t value);

void console_write_string(uint8_t* string);

const uint8_t test_Level_low = 0;
const uint8_t test_Level_high = 1;

const uint8_t test_Timer_timer1 = 1;
const uint8_t test_Timer_timer2 = 2;
const uint8_t test_Timer_timer3 = 3;
const uint8_t test_Timer_timer4 = 4;

struct test_Pwm
{
    uint32_t freq;
};

struct test_Cpu
{
    uint32_t osc;
    struct test_Pwm pwm;
};

struct test_Gpu
{
    struct test_Pwm* pwm;
};

uint8_t test_u8_data = 123;

uint8_t test_my_timer = 1;

struct test_Cpu test_cpu3 = {123, {456}};

void main()
{
    test_expression();
    test_statement();
}

void test_expression()
{
    console_write_string("============ test expression ============\n");
    test_test_unary();
    test_test_binary();
    test_test_conversion();
    test_test_increment();
    test_test_decrement();
    test_test_scope();
    test_test_initializer();
    test_test_member_access();
    test_test_subscripting();
    test_test_others();
}

void test_test_unary()
{
    global_assert(1, "constant true should equal true\n");
    global_assert(!0, "constant !false should equal true\n");
    int32_t v1 = +1;
    global_assert(v1 == 1, "unary (+), +1 should equal 1\n");
    int32_t v2 = -v1;
    global_assert(v2 == -1, "unary (-), -1 should equal -1\n");
    uint32_t v3 = ~1;
    global_assert(v3 == 4294967294, "unary (~), u32 ~1 should equal 4294967294\n");
    int32_t v4 = ~1;
    global_assert(v4 == -2, "unary (~), i32 ~1 should equal -2\n");
    uint8_t v5 = 0;
    global_assert(!v5, "unary (!), !false should equal true\n");
    int32_t* v6 = &v1;
    int32_t v7 = *v6;
    global_assert(v7 == 1, "unary (*), *variable should equal 1\n");
}

void test_test_increment()
{
    int32_t v = 0;
    v++;
    global_assert(v == 1, "increment 1++ should equal 2\n");
}

void test_test_decrement()
{
    int32_t v = 0;
    v--;
    global_assert(v == -1, "increment 1++ should equal 2\n");
}

void test_test_binary()
{
    int32_t v = 0;
    v = 5;
    global_assert(v == 5, "assign v to 5\n");
    v += 5;
    global_assert(v == 10, "v plus assign 5 should equal 10\n");
    v -= 2;
    global_assert(v == 8, "v minus assign 2 should equal 8\n");
    v *= 2;
    global_assert(v == 16, "v multi assign 2 should equal 16\n");
    v /= 4;
    global_assert(v == 4, "v divide assign 4 should equal 4\n");
    v %= 3;
    global_assert(v == 1, "v rem assign 4 should equal 1\n");
    v <<= 2;
    global_assert(v == 4, "v left shift assign 2 should equal 4\n");
    v >>= 1;
    global_assert(v == 2, "v right shift assign 1 should equal 2\n");
    v |= 1;
    global_assert(v == 3, "v or assign 1 should equal 3\n");
    v ^= 8;
    global_assert(v == 11, "v xor assign 8 should equal 11\n");
    v &= 6;
    global_assert(v == 2, "v and assign 6 should equal 2\n");
    global_assert((11 | 6) == 15, "11 | 6 should equal 15\n");
    global_assert((11 ^ 6) == 13, "11 ^ 6 should equal 13\n");
    global_assert((11 & 6) == 2, "11 & 6 should equal 2\n");
    global_assert(9 == 9, "9 == 9 should be true\n");
    global_assert(9 != 6, "9 != 6 should be true\n");
    global_assert(6 < 9, "6 < 9 should be true\n");
    global_assert(9 <= 9, "9 <= 9 should be true\n");
    global_assert(9 > 6, "9 > 6 should be true\n");
    global_assert(9 >= 9, "9 >= 9 should be true\n");
    global_assert((1 << 2) == 4, "1 << 2 should equal 4\n");
    global_assert((4 >> 1) == 2, "4 >> 1 should equal 2\n");
    global_assert(9 + 9 == 18, "9 + 9 should equal 18\n");
    global_assert(9 - 6 == 3, "9 - 6 should equal 3\n");
    global_assert(9 * 9 == 81, "9 * 9 should equal 81\n");
    global_assert(9 / 3 == 3, "9 / 3 should equal 3\n");
    global_assert(9 % 6 == 3, "9 % 6 should equal 3\n");
    global_assert((1 || 0) == 1, "true || false should be true\n");
    global_assert((1 && 0) == 0, "true && false should be false\n");
}

void test_test_others()
{
    global_assert((8 - 3) * 5 == 25, "(8 - 3) * 5 should equal 25\n");
    global_assert('a' == 97, "'a' should equal 97\n");
}

void test_test_initializer()
{
    uint8_t numbers[5] = {1, 2, 3, 4, 5};
    struct test_Cpu cpu = {123, {456}};
    struct test_Cpu cpus[2] = {{123, {456}}, {456, {789}}};
}

void test_test_subscripting()
{
    uint8_t numbers[5] = {1, 2, 3, 4, 5};
    global_assert(numbers[2] == 3, "numbers[2] should equal 3\n");
    numbers[2] = 9;
    global_assert(numbers[2] == 9, "numbers[2] should equal 9\n");
    uint8_t* array = numbers;
    array[2] = 8;
    global_assert(array[2] == 8, "array[2] should equal 9\n");
    uint8_t* array2 = array;
    array2[3] = 9;
    global_assert(array[3] == 9, "array2[3] should equal 9\n");
}

void test_test_scope()
{
    global_assert(test_u8_data == 123, "package var u8_data should equal 123\n");
    {
        global_assert(test_u8_data == 123, "package var u8_data should equal 123\n");
        uint8_t test_u8_data = 99;
        global_assert(test_u8_data == 99, "local var u8_data should equal 99\n");
        test_u8_data = 100;
    }
    global_assert(test_u8_data == 123, "package var u8_data should equal 123\n");
}

void test_test_member_access()
{
    struct test_Cpu cpu = {123, {456}};
    struct test_Cpu* cpu1 = &cpu;
    struct test_Cpu* cpu_array[2];
    cpu_array[0] = &cpu;
    cpu_array[1] = cpu1;
    global_assert(cpu.osc == 123, "cpu.osc should equal 123\n");
    global_assert(cpu1->osc == 123, "cpu1.osc should equal 123\n");
    global_assert(cpu_array[0]->osc == 123, "cpu_array[0].osc should equal 123\n");
    global_assert(cpu.pwm.freq == 456, "cpu.pwm.freq should equal 456\n");
    global_assert(cpu1->pwm.freq == 456, "cpu1.pwm.freq should equal 456\n");
    global_assert(cpu_array[0]->pwm.freq == 456, "cpu_array[0].pwm.freq should equal 456\n");
    global_assert(test_cpu3.osc == 123, "cpu3.osc should equal 123\n");
    global_assert(test_cpu3.osc == 123, "test.cpu3.osc should equal 123\n");
    global_assert(test_cpu3.pwm.freq == 456, "test.cpu3.pwm.freq should equal 456\n");
    struct test_Pwm pwm = {456};
    struct test_Gpu gpu;
    gpu.pwm = &pwm;
    global_assert(gpu.pwm->freq == 456, "gpu.pwm.freq should equal 456\n");
    struct test_Gpu* gpu1 = &gpu;
    global_assert(gpu1->pwm->freq == 456, "gpu1.pwm.freq should equal 456\n");
    struct test_Gpu* gpu_array[1];
    gpu_array[0] = &gpu;
    global_assert(gpu_array[0]->pwm->freq == 456, "gpu_array[0].pwm.freq should equal 456\n");
    global_assert(1 == 1, "enum Timer.timer1 should equal 1\n");
    global_assert(2 == 2, "enum Timer.timer2 should equal 2\n");
}

void test_test_conversion()
{
    int32_t a0 = 65856;
    int16_t a1 = ((int16_t)(a0));
    global_assert(a1 == 320, "convert i32 to i16, should equal 320\n");
    int8_t a2 = ((int8_t)(a1));
    global_assert(a2 == 64, "convert i16 to i8, should equal 64\n");
    float a3 = -3.14;
    int8_t a4 = ((int8_t)(a3));
    global_assert(a4 == -3, "convert f32 to i8, should equal -3\n");
    uint8_t a5 = ((uint8_t)(a3));
    global_assert(a5 == 253, "convert f32 to u8, should equal 253\n");
    float a6 = ((float)(a0));
    double a7 = ((double)(a0));
    float a8 = ((float)(a0));
    void* raw = &a3;
    int32_t* a9 = raw;
    int32_t a10 = *a9;
}

void global_assert(uint8_t expression, uint8_t* message)
{
    if (!expression)
    {
        console_write_string("assert failed:\n");
        console_write_string(message);
    }
}

void test_statement()
{
    console_write_string("============ test  statement ============\n");
    uint8_t a = 10;
    if (a >= 10)
    {
        global_assert(1, "a >= 10 should be true\n");
    }
    if (a < 10)
    {
        global_assert(0, "a < 10 shouldn't go here\n");
    }
    if (a > 100)
    {
        global_assert(0, "a > 100 shouldn't go here\n");
    }
    else 
    {
        global_assert(1, "(a > 100) else should be true\n");
    }
    uint8_t count = 0;
    for (;;)
    {
        count++;
        if (count == 3)
        {
            break;
        }
    }
    global_assert(count == 3, "count should equal 3\n");
    for (; count < 9;)
    {
        count++;
    }
    global_assert(count == 9, "count should equal 9\n");
    for (uint8_t i = 0; i < 10; i++)
    {
        if (i < 9)
        {
            continue;
        }
        count++;
    }
    global_assert(count == 10, "count should equal 10\n");
    uint8_t timer = 1;
    timer = 4;
    switch (timer)
    {
        case 1:
        case 2:
            global_assert(0, "case timer1 or timer2, shouldn't go here");
            break;
        case 3:
            global_assert(0, "case timer3, shouldn't go here");
            break;
        default:
            global_assert(1, "case timer4, should go here");
            count = 8;
    }
    global_assert(count == 8, "count should equal 8\n");
}

void console_write_bool(uint8_t value)
{
    if (value)
    {
        console_write_string("true");
    }
    else 
    {
        console_write_string("false");
    }
}

void console_write_i8(int8_t value)
{
    if (value < 0)
    {
        putchar(((int32_t)('-')));
        value = -value;
    }
    console_write_u64(((uint64_t)(value)));
}

void console_write_i16(int16_t value)
{
    if (value < 0)
    {
        putchar(((int32_t)('-')));
        value = -value;
    }
    console_write_u64(((uint64_t)(value)));
}

void console_write_i32(int32_t value)
{
    if (value < 0)
    {
        putchar(((int32_t)('-')));
        value = -value;
    }
    console_write_u64(((uint64_t)(value)));
}

void console_write_i64(int64_t value)
{
    if (value < 0)
    {
        putchar(((int32_t)('-')));
        value = -value;
    }
    console_write_u64(((uint64_t)(value)));
}

void console_write_u8(uint8_t value)
{
    console_write_u64(((uint64_t)(value)));
}

void console_write_u16(uint16_t value)
{
    console_write_u64(((uint64_t)(value)));
}

void console_write_u32(uint32_t value)
{
    console_write_u64(((uint64_t)(value)));
}

void console_write_u64(uint64_t value)
{
    uint8_t buffer[20];
    int8_t index = 0;
    for (; value != 0;)
    {
        uint8_t digit = ((uint8_t)(value / 10));
        buffer[index] = '0' + digit;
        index++;
        value = value / 10;
    }
    for (int8_t i = index - 1; i >= 0; i--)
    {
        putchar(((int32_t)(buffer[i])));
    }
}

void console_write_string(uint8_t* string)
{
    for (uint32_t i = 0; string[i] != 0; i++)
    {
        putchar(((int32_t)(string[i])));
    }
}

void test_Cpu_frequency()
{
    console_write_u32(123);
}

