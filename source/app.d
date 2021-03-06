import std.stdio;
import std.regex;
import std.file;
import std.path;
import std.string;
import std.array;
import std.algorithm;

immutable string CURRENT_INTERFACE_VERSION = "90002";

void writeResultsToFile(const string tocFileName, const string[] lines)
{
	auto f = File(tocFileName, "w");
	lines.each!(line => f.writeln(line));
}

void replaceInterfaceVersion()
{
	immutable string tocFile = getcwd.baseName ~ ".toc";

	if(tocFile.exists)
	{
		auto lines = tocFile.readText.lineSplitter();
		string[] outputLines;
		auto re = regex(r"(\w+)(\d+)","g");

		foreach(line; lines)
		{
			if(line.canFind("Interface:"))
			{
				immutable string replacedValue = replaceAll(line, re, CURRENT_INTERFACE_VERSION);
				outputLines ~= replacedValue;
			}
			else
			{
				outputLines ~= line;
			}
		}

		writeResultsToFile(tocFile, outputLines);
	}
}

void main(string[] arguments)
{
	replaceInterfaceVersion();
}
