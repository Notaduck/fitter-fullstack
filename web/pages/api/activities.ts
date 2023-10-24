import { withApiAuthRequired, getSession } from "@auth0/nextjs-auth0";

 export async function withApiAuthRequired(req, res) {
    const session = await getSession(req, res);

    // validate user can view courses
    res.send('1337');
}