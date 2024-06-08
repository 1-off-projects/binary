# loop menu
while True:
    # menu
    print("MENU:")
    print("1. to Binary")
    print("2. from binary")
    print("3. quit")
    choice = input("enter 1-3: ")
    # to
    if choice == '1':
        print("translating to binary")
        dec = int(input("decimal input: "))
        if dec == 0:
            dec == ("0b0")
        output = bin(dec)
        final = ("the equivelant is " + str(output[2:]))
        print((final))
        print("quit: y")
        print("don't: n")
        exit_prompt = input("answer: ")
        if exit_prompt == "y":
            exit
            break
        elif exit_prompt == "n":
            pass 
    # from
    elif choice == '2':
        print("translating from binary")
        binary_start = input("enter 1's and 0's: ")
        dec_end = int(binary_start, 2)
        dec_fin = ("the equivelant is " + str(dec_end))
        print(dec_fin)
        print("quit: y")
        print("don't: n")
        exit_prompt = input("answer: ")
        if exit_prompt == "y":
            exit
            break
        elif exit_prompt == "n":
            pass 
    # invalid & exit
    elif choice == '3':
        print("confirm: y")
        print("don't: n")
        exit_prompt = input("answer: ")
        if exit_prompt == "y":
            exit
            break
        elif exit_prompt == "n":
            pass 
    else:
        print("invalid, goodbye...")
        exit
        break