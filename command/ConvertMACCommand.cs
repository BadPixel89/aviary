//  C# code that converts macs, will port to go and simplify
//  probably just make it parse the mac into all types and print
//  later add the ability to specify just one type
//  specify clipboard behaviour - default don't use clipboard
  //  can have the behavoiur in the conf

public string Name { get; set; } = "mac";
private enum MACTYPE { NONE = 0, COLON, DASH, DOT, ERROR, ALL }

public void macConvert(string[] args)
{
    if (args.Length > 0)
    {
        string rawMac = GetRawMac(args[0]);

        if (ValidateRawMac(rawMac))
        {
            try
            {
                switch (args[1])
                {
                    case "raw":
                    case "-raw":
                    case "-r":
                        OutputMacs(rawMac,MACTYPE.NONE,args.Contains("-l"));
                        break;
                    case "colon":
                    case "-colon":
                    case "-c":
                        OutputMacs(rawMac, MACTYPE.COLON, args.Contains("-l"));
                        break;
                    case "dash":
                    case "-dash":
                    case "-d":
                        OutputMacs(rawMac, MACTYPE.DASH, args.Contains("-l"));
                        break;
                    case "dot":
                    case "-dot":
                    case "-D":
                        OutputMacs(rawMac, MACTYPE.DOT, args.Contains("-l"));
                        break;
                    case "all":
                    case "-a":
                        OutputMacs(rawMac, MACTYPE.ALL, args.Contains("-l"));
                        break;
                    default:
                        OutputMacs(rawMac, MACTYPE.ERROR, args.Contains("-l"));
                        break;
                }
            }
            catch (IndexOutOfRangeException)
            {
                OutputMacs(rawMac, MACTYPE.ERROR, args.Contains("-l"));
            }
        }
    }
    else
    {
        Logger.WriteError("must provide a mac address");
    }
}
private string GetRawMac(string macToConvert)
{
    switch (DetectFormat(macToConvert))
    {
        case MACTYPE.NONE:
            return macToConvert;
        case MACTYPE.COLON:
            return string.Join("", macToConvert.Split(':'));
        case MACTYPE.DASH:
            return string.Join("", macToConvert.Split('-'));
        case MACTYPE.DOT:
            return string.Join("", macToConvert.Split('.'));
        case MACTYPE.ERROR:
            Logger.WriteError("invalid mac entered");
            break;
        default:
            break;
    }
    return null;
}
private MACTYPE DetectFormat(string inputMac)
{
    if (string.IsNullOrEmpty(inputMac))
    {
        return MACTYPE.ERROR;
    }
    if (inputMac.Contains(':'))
    {
        return MACTYPE.COLON;
    }
    if (inputMac.Contains('-'))
    {
        return MACTYPE.DASH;
    }
    if (inputMac.Contains('.'))
    {
        return MACTYPE.DOT;
    }
    return MACTYPE.NONE;
}
private bool ValidateRawMac(string inputMac)
{
    if (inputMac.Length != 12)
    {
        return false;
    }
    long output;
    if (!long.TryParse(inputMac, System.Globalization.NumberStyles.HexNumber, null, out output))
    {
        return false;
    }

    return true;
}
private void OutputMacs(string macNoFormatting,MACTYPE requestedClipboardContents, bool lowerCase)
{
    Dictionary<MACTYPE, string> macs = new Dictionary<MACTYPE, string>();
    string output = macNoFormatting.ToUpper();

    macs.Add(MACTYPE.NONE, output);
    macs.Add(MACTYPE.COLON, string.Empty);
    macs.Add(MACTYPE.DASH, string.Empty);
    macs.Add(MACTYPE.DOT, string.Empty);
    macs.Add(MACTYPE.ALL, string.Empty);

    try
    {
        //this loop runs 12 times, using the counter -1 to index arrays so we can count elements and indexes
        for (int i = 1; i < 13; i++)
        {
            macs[MACTYPE.COLON] += output[i - 1];
            macs[MACTYPE.DASH] += output[i - 1];
            macs[MACTYPE.DOT] += output[i - 1];

            if (i == 12) { break; }

            if (i % 2 == 0)
            {
                macs[MACTYPE.COLON] += ":";
                macs[MACTYPE.DASH] += "-";
            }
            if (i % 4 == 0)
            {
                macs[MACTYPE.DOT] += ".";
            }
        }
        macs[MACTYPE.ALL] += 
            macs[MACTYPE.NONE] + "\r\n" +
            macs[MACTYPE.COLON] + "\r\n" +
            macs[MACTYPE.DASH] + "\r\n" +
            macs[MACTYPE.DOT] + "\r\n";
            
    }
    catch (Exception e)
    {
        Logger.WriteError(e.Message + "\r\n" + "error occurred during conversion");
        return;
    }

    string consoleOutput = string.Empty;
    consoleOutput +=
        "raw:".PadRight(10) + output + "\r\n" +
        "colon:".PadRight(10) + macs[MACTYPE.COLON] + "\r\n" +
        "dash:".PadRight(10) + macs[MACTYPE.DASH] + "\r\n" +
        "dot:".PadRight(10) + macs[MACTYPE.DOT];
    Logger.WriteStatus(consoleOutput);

    if(requestedClipboardContents != MACTYPE.ERROR)
    {
        if (lowerCase)
        {
            Clipboard.SetText(macs[requestedClipboardContents].ToLower());
            return;
        }
        Clipboard.SetText(macs[requestedClipboardContents]);
    }
    if(requestedClipboardContents == MACTYPE.ALL)
    {
        if (lowerCase)
        {
            Clipboard.SetText(macs[requestedClipboardContents].ToLower());
            return;
        }
        Clipboard.SetText(macs[requestedClipboardContents]);
        return;
    }

    return;
}
