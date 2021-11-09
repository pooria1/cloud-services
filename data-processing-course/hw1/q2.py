import pandas as pd 

data = pd.read_csv("players.csv") 
    
print("\n\n\n\n---------------PART A---------------\n\n")
print(f'first and last index of file is:\n{data.iloc[[0, -1]]}')

print("\n\n\n\n---------------PART C---------------\n\n")
print(f'minimum weight is {data["Weight"].min()} and maximum is {data["Weight"].max()}')
print(f'average of weight is {data["Weight"].mean()}')

print("\n\n\n\n---------------PART D---------------\n\n")
a = data.groupby(['Nationality']).size().reset_index(name='players_count')
print(f'minumum country players is:\n {a[a["players_count"]==a["players_count"].min()]} \nand maximum is:\n\n {a[a["players_count"]==a["players_count"].max()]}')

print("\n\n\n\n---------------PART E---------------\n\n")
print(f'players with growth < 3 and potential < 84 are:\n{data.query("Growth < 3 and Potential < 84")}')

print("\n\n\n\n---------------PART F---------------\n\n")
print(f'past report with position:\n{data.query("Growth < 3 and Potential < 84")[["Name", "Growth", "Potential", "Positions"]]}')
  
print("\n\n\n\n---------------PART H---------------\n\n")
a = data.query("ContractUntil < 2021 and NationalTeam == 'Not in team'")
print(f'number of players contract until 2021 and not in national team: {len(a)}')
print(f'those players are:\n{a}')

print("\n\n\n\n---------------PART I---------------\n\n")
a = data.query("Age < 24 and Potential < 84 and Club == 'Chelsea'")[["Name", "Age", "ValueEUR"]]
print(f'chelsea under 24 years old\'s value:\n{a}')

print("\n\n\n\n---------------PART J---------------\n\n")
a = data.query("Name == 'E. Hazard'")[["Name", "Club", "Positions","WageEUR"]]
print(f'hazard\'s info is:\n{a}')

