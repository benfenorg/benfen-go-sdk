import random


def random_user():
    data = ''
    for i in range(8):
        r1 = chr(random.randint(97, 122))
        data = data + r1
    data = data + "@aa.com"
    return data

def random_account(index):
    data = ["ff6372f3-d3b5-45cd-b826-34607b69acb3",
            "0d7e6d5e-f285-4fc5-a8c2-15aeb5d0648a",
            "3e632896-249a-4987-9be5-86cdc1414093",
            "48040b9b-f095-47e5-8ad0-12627d7247a4",
            "c3ab44e1-a485-47b6-b7f3-512e977f5bc5",
            "6b183021-4d5e-4be1-aa97-ea0a08be5b45", "2461a727-71bd-4aa8-973d-beb0508d79ad",
            "f39d05d3-098a-4663-8c17-8f923efec32e",
            "eda57e8a-3f9c-4506-a59e-2daa3973d9f8", "83c4b555-2e15-4a33-9dc1-54bcc654b32c",
            "87df1616-e8a2-44db-aa63-3b2a5ad3d6c0", "e3ff8ba3-cd5c-47c4-9492-5e68bc9388f2",
            "f831ed54-dfc6-46a4-a4d8-f19ef2a73c20", "ffc9b09b-9669-4e53-8b86-f09f6d465efd",
            "8dc796fd-bc9a-4c03-89b0-c0b45e2e8706", "d08e852f-7970-4ff9-a198-f482bbe64eab",
            "45837788-1ea1-4a77-aa0c-8bbb5bc866a7", "45116689-aa28-4dd6-9e33-a80caeb9a2b6",
            "4dabc692-9deb-44f9-a9ee-e0e9a0c19b29", "43b31693-a61e-43ac-8ba7-a01848f7e04e",
            "ff62217e-8304-471b-a858-28931f4f99e9",
            "553cb056-a13f-444b-b970-d600aa2b0824", "789e0629-88f5-41af-a440-6690ac8bee7b"]
    return data[index % len(data)]
