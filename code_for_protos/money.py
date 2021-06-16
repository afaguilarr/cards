from python_proto_code.protos.person_pb2 import Money, DayOfWeek, Person
from python_proto_code.protos.date_pb2 import Date

money = Money(
    currency_code="USD",
    integral_amount=12,
    decimal_amount=12
)

print(money)

print(money.SerializeToString())

with open("code_for_protos/money.bin", "wb") as file:
    file.write(money.SerializeToString())

with open("code_for_protos/money.bin", "rb") as file:
    read_money = Money.FromString(file.read())

print(read_money)

day = DayOfWeek(
    day=DayOfWeek.SUNDAY
)

print(day.day == DayOfWeek.SUNDAY)
print(day.day == 7)
print(day.day == "SUNDAY")

print(day)

print(DayOfWeek.SUNDAY)
print(day.day)

person = Person(
    birthday_date=Date(
        day=21,
        month=12,
        year=2012
    ),
    phone_numbers=["3433277", "3427224", "5888691", "4428924"],
    dates=[
        Date(
            day=21,
            month=12,
            year=2012
        ),
        Date(
            day=21,
            month=12,
            year=2021
        ),
    ]
)

print(person)

person.phone_numbers.append("1234567")

person.dates.append(
    Date(
        day=8,
        month=8,
        year=2018
    ),
)
new_date = person.dates.add()
new_date.day = 9
new_date.month = 9
new_date.year = 2019

person.phone_numbers.extend(["7654321"])

print(person)
