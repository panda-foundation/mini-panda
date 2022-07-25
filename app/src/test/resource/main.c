#include <stdint.h>

struct test_Cpu;

void main();

void test_expression();

void test_test_unary();

void test_test_increment();

void test_test_decrement();

void test_test_binary();

void test_test_others();

void test_test_initializer();

void global_assert(uint8_t expression, uint8_t* message);

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

uint8_t test_Level_low = 0;
uint8_t test_Level_high = 1;

uint8_t test_Timer_timer1 = 1;
uint8_t test_Timer_timer2 = 2;
uint8_t test_Timer_timer3 = 3;
uint8_t test_Timer_timer4 = 4;

struct test_Cpu
{
    uint32_t osc;
};

void main(){
    test_expression();
}

void test_expression(){
    console_write_string("============ test expression ============\n");
    test_test_unary();
    test_test_binary();
    test_test_increment();
    test_test_decrement();
    test_test_initializer();
    test_test_others();
}

void test_test_unary(){
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

void test_test_increment(){
    int32_t v = 0;
    v++;
    global_assert(v == 1, "increment 1++ should equal 2\n");
}

void test_test_decrement(){
    int32_t v = 0;
    v--;
    global_assert(v == -1, "increment 1++ should equal 2\n");
}

void test_test_binary(){
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

void test_test_others(){
    global_assert((8 - 3) * 5 == 25, "(8 - 3) * 5 should equal 25\n");
    global_assert('a' == 97, "'a' should equal 97\n");
}

void test_test_initializer(){
    uint8_t numbers[5] = {1, 2, 3, 4, 5};
}

void global_assert(uint8_t expression, uint8_t* message){
    if (!expression)
    {
        console_write_string("assert failed:\n");
        console_write_string(message);
        console_write_u8('\n');
    }
}

void console_write_bool(uint8_t value){
    if (value)
    {
        console_write_string("true");
    }
    else 
    {
        console_write_string("false");
    }
}

void console_write_i8(int8_t value){
    if (value < 0)
    {
        putchar(((int32_t)('-')));
        value = -value;
    }
    console_write_u64(((uint64_t)(value)));
}

void console_write_i16(int16_t value){
    if (value < 0)
    {
        putchar(((int32_t)('-')));
        value = -value;
    }
    console_write_u64(((uint64_t)(value)));
}

void console_write_i32(int32_t value){
    if (value < 0)
    {
        putchar(((int32_t)('-')));
        value = -value;
    }
    console_write_u64(((uint64_t)(value)));
}

void console_write_i64(int64_t value){
    if (value < 0)
    {
        putchar(((int32_t)('-')));
        value = -value;
    }
    console_write_u64(((uint64_t)(value)));
}

void console_write_u8(uint8_t value){
    console_write_u64(((uint64_t)(value)));
}

void console_write_u16(uint16_t value){
    console_write_u64(((uint64_t)(value)));
}

void console_write_u32(uint32_t value){
    console_write_u64(((uint64_t)(value)));
}

void console_write_u64(uint64_t value){
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

void console_write_string(uint8_t* string){
    for (uint32_t i = 0; string[i] != 0; i++)
    {
        putchar(((int32_t)(string[i])));
    }
}

void test_Cpu_frequency(){
    console_write_u32(123);
}

