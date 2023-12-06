
import { getAccessToken, getSession } from '@auth0/nextjs-auth0'
import axios, { AxiosHeaders, AxiosRequestConfig, AxiosRequestHeaders, AxiosResponseHeaders } from 'axios'
import formidable, { IncomingForm } from 'formidable'
import multer from 'multer'
import type { NextApiRequest, NextApiResponse } from 'next'

type ResponseData = {
  message: string
}

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<ResponseData>
) {

  if (req.method === 'GET') {

    const activityId = req?.query?.activityId;

    if (!activityId) {
      return res.status(400).json({ error: 'Missing activityId parameter' });
    }

    try {


      const { idToken } = await getSession(req, res)
      const headers: Partial<AxiosHeaders> = {
        Authorization: `Bearer ${idToken}`
      }

      const activities = await axios.get(`http://fitter-go:3030/activity?id=${req.query.activityId}`, {
        headers
      })



      res.status(activities.status).json(activities.data)
    } catch (err) {
      console.error(err)
      res.status(500).json(err);
    }




  }

  if (req.method === "POST") {


    // const storage = multer.memoryStorage();
    // const upload = multer({ storage: storage });

    // const { idToken } = await getSession(req, res)

    // const headers: Partial<AxiosHeaders> = {
    //   Authorization: `Bearer ${idToken}`
    // }

    // console.log(req)


    // const formData = new FormData();
    // req.files.forEach((file) => {
    //   formData.append('files', file.buffer, { filename: file.originalname });
    // });
    try {


      let status = 200,
        resultBody = { status: 'ok', message: 'Files were uploaded successfully' };

      /* Get files using formidable */
      const files = await new Promise<ProcessedFiles | undefined>((resolve, reject) => {
        const form = new IncomingForm({ keepExtensions: true });
        const files: ProcessedFiles = [];
        form.on('fils', function (field, file) {
          files.push([field, file]);
        })
        form.on('end', () => resolve(files));
        form.on('error', err => reject(err));
        form.parse(req, () => {
          //
        });
      }).catch(e => {
        console.log(e);
        status = 500;
        resultBody = {
          status: 'fail', message: 'Upload error'
        }
      });

      console.log('files', files)

      // console.log(req.)
      // console.log(re)

      // const formData = new FormData();
      // req.files.forEach((file) => {
      //   formData.append('files', file.buffer, { filename: file.originalname });
      // });

      // const response = await axios.post(apiUrl, formData, {
      //   headers: {
      //     'Content-Type': 'multipart/form-data',
      //     ...formData.getHeaders(),
      //   },
      // });

      // // Send the response from the other API back to the client
      // res.status(res.status).json(response.data);

    } catch (error) {
      console.error('Error:', error);
      res.status(500).json(error);
    }

  }

  return res.status(400).json({ error: 'method not supported' })
}


/* Don't miss that! */
export const config = {
  api: {
    bodyParser: false,
  }
};

type ProcessedFiles = Array<[string, File]>;