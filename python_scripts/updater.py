import asyncio

import aiohttp


async def processing(client_session):
    response = await client_session.get('http://grademylook.com/images')
    response = await response.read()
    print(response)
    await asyncio.sleep(2)
    await do_request(client_session)


async def do_request(client_session):
    await processing(client_session)


loop = asyncio.get_event_loop()

try:
    client_session = aiohttp.ClientSession(connector=aiohttp.TCPConnector(ssl=False))
    loop.run_until_complete(do_request(client_session))
    loop.run_forever()

finally:

    print('closing event loop')
    loop.close()
