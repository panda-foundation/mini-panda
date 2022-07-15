#include <stdint.h>

void main();

void test_expression();

void test_test_unary();

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

void main(){
    test_expression();
}

void test_expression(){
    console_write_string("============ test expression ============\n");
    test_test_unary();
}

void test_test_unary(){
    global_assert(1, "constant true\n");
    global_assert(!0, "constant false\n");
    int32_t v1 = +1;
    global_assert(v1 == 1, "unary (+), v1 should equal 1\n");
    int32_t v2 = -v1;
    global_assert(v2 == -1, "unary (-), v2 should equal -1\n");
    uint32_t v3 = ~1;
    global_assert(v3 == 4294967294, "unary (~), v3 should equal 4294967294\n");
    int32_t v4 = ~1;
    global_assert(v4 == -2, "unary (~), v4 should equal -2\n");
    uint8_t v5 = 0;
    global_assert(!v5, "unary (!), v5 should equal false\n");
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
        putchar(((int32_t)'-'));
        value = -value;
    }
    console_write_u64(((uint64_t)value));
}

void console_write_i16(int16_t value){
    if (value < 0)
    {
        putchar(((int32_t)'-'));
        value = -value;
    }
    console_write_u64(((uint64_t)value));
}

void console_write_i32(int32_t value){
    if (value < 0)
    {
        putchar(((int32_t)'-'));
        value = -value;
    }
    console_write_u64(((uint64_t)value));
}

void console_write_i64(int64_t value){
    if (value < 0)
    {
        putchar(((int32_t)'-'));
        value = -value;
    }
    console_write_u64(((uint64_t)value));
}

void console_write_u8(uint8_t value){
    console_write_u64(((uint64_t)value));
}

void console_write_u16(uint16_t value){
    console_write_u64(((uint64_t)value));
}

void console_write_u32(uint32_t value){
    console_write_u64(((uint64_t)value));
}

void console_write_u64(uint64_t value){
    uint8_t buffer[20];
    int8_t index = 0;
    for (; value != 0;)
    {
        uint8_t digit = ((uint8_t)value / 10);
        buffer[index] = '0' + digit;
        index++;
        value = value / 10;
    }
    for (int8_t i = index - 1; i >= 0; i--)
    {
        putchar(((int32_t)buffer[i]));
    }
}

void console_write_string(uint8_t* string){
    for (uint32_t i = 0; string[i] != 0; i++)
    {
        putchar(((int32_t)string[i]));
    }
}

