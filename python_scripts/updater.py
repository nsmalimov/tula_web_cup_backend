import asyncio

import aiohttp


async def processing(client_session):
    try:
        response = await client_session.get('https://grademylook.com/backend/images_update_urls')
        response = await response.read()
    except Exception as e:
        print(e)

    await asyncio.sleep(1 * 60 * 60)
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
