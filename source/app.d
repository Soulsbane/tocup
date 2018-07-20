import std.stdio;
import std.regex;
import std.file;
import std.path;
import std.string;
import std.array;
import std.algorithm;

immutable string CURRENT_INTERFACE_VERSION = "80000";

void writeResultsToFile(string[] lines)
{
	foreach(line; lines)
	{
		writeln(line);
	}
}

void replaceInterfaceVersion()
{
	immutable string tocFile = getcwd.baseName ~ ".toc";

	if(tocFile.exists)
	{
		auto lines = tocFile.readText.lineSplitter();
		string[] outputLines;

		foreach(line; lines)
		{
			if(line.canFind("Interface:"))
			{
				auto re = regex(r"(\w+)(\d+)","g");
				immutable string replacedValue = replaceAll(line, re, CURRENT_INTERFACE_VERSION);

				outputLines ~= replacedValue;
			}
			else
			{
				outputLines ~= line;
			}
		}

		writeResultsToFile(outputLines);
	}
}

void main(string[] arguments)
{
	replaceInterfaceVersion();
}
