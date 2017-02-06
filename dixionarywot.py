#! /usr/bin/python3


import discord
import json

client = discord.Client()
ovners = ['201741426038538242', '194535808940769280', '206816710886555659', '205377718450257920', '109710323094683648']
sign = '~'
nocorrect = list()

dix = dict()

with open('nocorrect.json', 'r') as f:
    nocorrect = json.load(f)

with open('dixionary.json', 'r') as f:
    dix = json.load(f)

state = True

@client.event
async def on_ready():
    print('Logged in as')
    print(client.user)
    print(client.user.id)
    print('--------------')

    
@client.event
async def on_message(message):
    global state
    global nocorrect
    if not message.author.bot:
        if state and not message.author.id in nocorrect:
            correction = ""
            words = message.content.split(' ')
            for word in words:
                for key in dix:
                    if key.lower() == word.lower():
                        correction = "{} *{}".format(correction, dix[key])
            if not correction == "":
                await client.send_message(message.channel, "{}".format(correction))
                
    if message.content.startswith(sign) and message.author.id in ovners:
        if message.content.startswith('{}off'.format(sign)):
            state = False
            await client.send_message(message.channel, "```Corrections off```")
            
        if message.content.startswith('{}on'.format(sign)):
            state = True
            await client.send_message(message.channel, "```Corrections on```")
            
        if message.content.startswith('{}add'.format(sign)):
            splitmsg = message.content.split(' ')
            dix[splitmsg[1]] = splitmsg[2]
            with open('dixionary.json', 'w') as f:
                json.dump(dix, f)
            await client.send_message(message.channel, "```Now correcting {} to {}```".format(splitmsg[1],splitmsg[2]))
            
    elif message.content.startswith('{}off'.format(sign)):
        nocorrect.append(message.author.id)
        with open('nocorrect.json', 'w') as f:
            json.dump(nocorrect, f)
        await client.send_message(message.channel, "```OK, {}, I vill no longer correct you```".format(message.author.name))
        
    elif message.content.startswith('{}on'.format(sign)):
        idx = nocorrect.index(message.author.id)
        nocorrect.pop(idx)
        with open('nocorrect.json', 'w') as f:
            json.dump(nocorrect, f)
        await client.send_message(message.channel, "```OK, {}, I vill correct you again```".format(message.author.name))
        
    if message.content == ('{}inwitessl'.format(sign)):
        msg = 'http://discord.me/ScammerSubLounge {0.mention}'.format(message)
        await client.send_message(message.channel, msg)
        

    if message.content == ('{}inwitedixionary'.format(sign)):
        msg = 'http://discord.me/Dixionary {0.mention}'.format(message)
        await client.send_message(message.channel, msg)
        
        


                       
@client.event
async def on_member_join(member):
    server = member.server
    fmt = 'Velcome {0.mention} to {1.name}!'
    await client.send_message(server, fmt.format(member, server))
    return


        
           
client.run('')
