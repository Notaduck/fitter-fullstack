import axios from 'axios';
import React, { useState } from 'react';
import { useDropzone } from 'react-dropzone';

const FileUpload = () => {
    const [uploadedFiles, setUploadedFiles] = useState<File[]>([]);
    const { getRootProps, getInputProps } = useDropzone({
        accept: {
            'application/octet-stream': [".fit"]
        },
        onDrop: async (acceptedFiles: File[]) => {

            setUploadedFiles(acceptedFiles);

            const formData = new FormData()

            acceptedFiles.forEach(file => {
                formData.append('file', file)
            });

            try {

                const res = await axios.post('/api/activity', {
                    data: formData
                    header
                })

                console.log('res', res)

            } catch (err) { console.error(err) }




            console.log(acceptedFiles)
            // axios.post('/api/activity', )
        },

    });
    //TO DO : Customize and Style this Drag and Drop to Upload box as you want🧑‍💻😊
    return (
        <div className='flex-1 flex flex-col items-center bg-neutral-50 text-[#bdbdbd] transition-[border] duration-[0.24s] ease-[ease-in-out] p-5 rounded-sm border-2 border-zinc-100 hober:border-zinc-400 border-dashed' {...getRootProps()}>
            <input {...getInputProps()} />
            <p>Drag and drop files here or click to add <strong> new activities</strong>.</p>
            <progress className="progress w-56"></progress>
            <ul>
                {uploadedFiles.map((file) => (
                    <li key={file.name}>{file.name}</li>
                ))}
            </ul>
        </div>
    );
};
export default FileUpload;