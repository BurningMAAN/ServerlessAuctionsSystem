import CsvDownload from 'react-json-to-csv'

export const generateData = (data) => {
    console.log("invoked")
    return <CsvDownload data={data} filename="ataskaita.csv"></CsvDownload>
}